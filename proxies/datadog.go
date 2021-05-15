package proxies

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

type DataDogProxy struct {
	ctx       context.Context
	apiClient *datadog.APIClient
}

func GetDataDogProxy() *DataDogProxy {
	ddp := DataDogProxy{}

	ddp.ctx = datadog.NewDefaultContext(context.Background())

	configuration := datadog.NewConfiguration()
	ddp.apiClient = datadog.NewAPIClient(configuration)

	return &ddp
}

func (ddp *DataDogProxy) GetEvents(
	source string,
	startTime int64,
	endTime int64,
	eventPriority string) datadog.EventListResponse {
	priority := datadog.EventPriority(eventPriority)
	sources := source
	// tags := ""
	unaggregated := true
	excludeAggregate := true
	// page := int32(56)
	optionalParams := datadog.ListEventsOptionalParameters{
		Priority: &priority,
		Sources:  &sources,
		// Tags:             &tags,
		Unaggregated:     &unaggregated,
		ExcludeAggregate: &excludeAggregate,
		// Page:             &page,
	}

	resp, r, err := ddp.apiClient.EventsApi.ListEvents(ddp.ctx, startTime, endTime, optionalParams)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	return resp
}
