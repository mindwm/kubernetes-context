package neo4j

import (
	"context"
	"errors"
	"fmt"
	"function/entity"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repository struct {
	driver neo4j.DriverWithContext
}

func NewRepository(driver neo4j.DriverWithContext) *Repository {
	return &Repository{driver: driver}
}

func (r Repository) GetIODocumentByID(ctx context.Context, ioDocumentID int) (entity.IODocument, error) {
	const op = "Neo4j.Repository.GetIODocumentByID"

	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := "MATCH (n) WHERE ID(n) = $ioDocumentID RETURN n"
	result, err := session.Run(ctx, query, map[string]interface{}{"ioDocumentID": ioDocumentID})
	if err != nil {
		return entity.IODocument{}, fmt.Errorf("%s:%v", op, err)
	}

	var ioDoc entity.IODocument
	if result.Next(ctx) {
		record := result.Record()
		ioDoc, err = convertRecordToIODocument(record)
		if err != nil {
			return entity.IODocument{}, fmt.Errorf("%s:%v", op, err)
		}
	}

	return ioDoc, nil
}

func convertRecordToIODocument(record *neo4j.Record) (entity.IODocument, error) {
	node, ok := record.Get("n")
	if !ok {
		return entity.IODocument{}, errors.New("no node found in record")
	}

	ioDoc := entity.IODocument{}
	props := node.(neo4j.Node).Props

	if elementID, ok := props["elementId"].(int); ok {
		ioDoc.ElementID = elementID
	}
	if id, ok := props["id"].(int); ok {
		ioDoc.ID = id
	}
	if output, ok := props["output"].(string); ok {
		ioDoc.Output = output
	}
	if ps1, ok := props["ps1"].(string); ok {
		ioDoc.PS1 = ps1
	}
	if time, ok := props["time"].(float64); ok {
		ioDoc.Time = time
	}
	if userInput, ok := props["user_input"].(string); ok {
		ioDoc.UserInput = userInput
	}
	if uuid, ok := props["uuid"].(string); ok {
		ioDoc.UUID = uuid
	}

	return ioDoc, nil
}

func (r Repository) AddInterfacesToNode(ctx context.Context, nodeID int, iface entity.Interface) error {
	const op = "Neo4j.Repository.AddInterfaceToNode"

	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		query := `
            MATCH (n)
            WHERE id(n) = $nodeId
            CREATE (n)-[:HAS_INTERFACE]->(i:Interface {
                name: $name,
                mtu: $mtu,
                rx_pkt: $rx_pkt,
                rx_bts: $rx_bts,
                tx_pkt: $tx_pkt,
                tx_bts: $tx_bts
            })
            RETURN id(i)
        `
		params := map[string]interface{}{
			"nodeId": nodeID,
			"name":   iface.Name,
			"mtu":    iface.MTU,
			"rx_pkt": iface.RxPkt,
			"rx_bts": iface.RxBytes,
			"tx_pkt": iface.TxPkt,
			"tx_bts": iface.TxBytes,
		}

		result, err := tx.Run(ctx, query, params)
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}

	return nil
}
