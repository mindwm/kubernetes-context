package function

import (
	"context"
	"fmt"
	"function/db"
	neo4jRepo "function/db/repository/neo4j"
	"function/entity"
	"function/parser"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"strconv"
)

type Repository interface {
	GetIODocumentByID(ctx context.Context, ioDocumentID int) (entity.IODocument, error)
	AddInterfaceToNode(ctx context.Context, nodeID int, iface entity.Interface) error
	AddPodToNode(ctx context.Context, nodeID int, pod entity.Pod) error
	AddServiceToNode(ctx context.Context, nodeID int, service entity.Service) error
	AddNamespaceToNode(ctx context.Context, nodeID int, namespace entity.NameSpace) error
}

// Handle an HTTP Request.
func Handle(ctx context.Context, event cloudevents.Event) error {
	var cloudEventData entity.CloudEventData
	if err := event.DataAs(&cloudEventData); err != nil {
		fmt.Println(err)
		return nil
	}

	if isRelationship(cloudEventData) {
		nodeID, err := strconv.Atoi(cloudEventData.Payload.End.ID)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		driver := db.InitNeo4j(ctx)
		defer driver.Close(ctx)
		var repository Repository = neo4jRepo.NewRepository(driver)

		ioDoc, err := repository.GetIODocumentByID(ctx, nodeID)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(ioDoc.Output)

		if parser.IsIfConfig(ioDoc.UserInput) {
			networkInfo, err := parser.ParseIfConfigOutput(ioDoc.Output)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			fmt.Println(len(networkInfo))

			for _, iface := range networkInfo {
				err = repository.AddInterfaceToNode(ctx, nodeID, iface)
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
		} else if parser.IsKubectlGetPods(ioDoc.UserInput) {
			pods, err := parser.ParseKubectlGetPodsOutput(ioDoc.Output)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			fmt.Println(len(pods))

			for _, pod := range pods {
				err = repository.AddPodToNode(ctx, nodeID, pod)
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
		} else if parser.IsKubectlGetServices(ioDoc.UserInput) {
			services, err := parser.ParseKubectlGetSvcOutput(ioDoc.Output)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			for _, service := range services {
				err = repository.AddServiceToNode(ctx, nodeID, service)
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
		} else if parser.IsKubectlGetNamespaces(ioDoc.UserInput) {
			namespaces, err := parser.ParseKubectlGetNsOutput(ioDoc.Output)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			fmt.Println(len(namespaces))

			for _, namespace := range namespaces {
				err = repository.AddNamespaceToNode(ctx, nodeID, namespace)
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
		}
	}

	return nil

}

func isRelationship(eventData entity.CloudEventData) bool {
	return eventData.Payload.Type == "relationship" && eventData.Payload.Label == "HAS_IO_DOCUMENT"
}
