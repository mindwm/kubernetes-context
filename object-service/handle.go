package function

import (
	"context"
	"fmt"
	"function/db"
	"function/db/repository/neo4j"
	"function/entity"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"net/http"
	"os"
)

type Repository interface {
	Save(ctx context.Context, host entity.Host) error
}

// Handle an HTTP Request.
func Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	event, err := cloudevents.NewEventFromHTTPRequest(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get new event from request")
		return
	}

	var request entity.CloudEventRequest
	if err := event.DataAs(&request); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode cloud event data: %v", err)
		return
	}

	host := entity.Host{
		Hostname:   r.RemoteAddr,
		Interfaces: request.NetworkInfo,
	}

	driver := db.InitNeo4j(ctx)
	defer driver.Close(ctx)
	var hostRepository Repository = neo4j.NewRepository(driver)

	if err := hostRepository.Save(ctx, host); err != nil {
		fmt.Fprintf(os.Stderr, "failed to save host data: %v", err)
		return
	}

	fmt.Fprintln(w, "data saved successfully")
}
