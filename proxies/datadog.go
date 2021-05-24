package proxies

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/vivasaayi/cloudrover/utililties"
)

type DataDogProxy struct {
	ctx       context.Context
	apiClient *datadog.APIClient
	apiKey    string
	appKey    string
}

func GetDataDogProxy() *DataDogProxy {
	ddp := DataDogProxy{}

	ddp.ctx = datadog.NewDefaultContext(context.Background())

	configuration := datadog.NewConfiguration()
	ddp.apiClient = datadog.NewAPIClient(configuration)

	ddp.apiKey = utililties.GetStringEnvVar("DD_API_KEY", "", true)
	ddp.appKey = utililties.GetStringEnvVar("DD_APP_KEY", "", true)

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

func (ddp *DataDogProxy) GetMonitors() []datadog.Monitor {
	groupStates := "all" // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)
	// name := "name_example"               // string | A string to filter monitors by name. (optional)
	// tags := "tags_example"               // string | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, `host:host0`. (optional)
	// monitorTags := "monitorTags_example" // string | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, `service:my-app`. (optional)
	withDowntimes := true  // bool | If this argument is set to true, then the returned data includes all current downtimes for each monitor. (optional)
	idOffset := int64(789) // int64 | Monitor ID offset. (optional)
	page := int64(789)     // int64 | The page to start paginating from. If this argument is not specified, the request returns all monitors without pagination. (optional)
	pageSize := int32(56)  // int32 | The number of monitors to return per page. If the page argument is not specified, the default behavior returns all monitors without a `page_size` limit. However, if page is specified and `page_size` is not, the argument defaults to 100. (optional)
	optionalParams := datadog.ListMonitorsOptionalParameters{
		GroupStates: &groupStates,
		// Name:          &name,
		// Tags:          &tags,
		// MonitorTags:   &monitorTags,
		WithDowntimes: &withDowntimes,
		IdOffset:      &idOffset,
		Page:          &page,
		PageSize:      &pageSize,
	}

	resp, r, err := ddp.apiClient.MonitorsApi.ListMonitors(ddp.ctx, optionalParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ListMonitors`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	return resp
}

func (ddp *DataDogProxy) SearchMonitors() DDMonitorSearchResponse {
	client := &http.Client{}

	req, err := http.NewRequest("GET",
		`https://api.datadoghq.com/api/v1/monitor/search?query=status:alert&per_page=1000`,
		nil,
	)

	if err != nil {
		fmt.Println("Error occured when retrieving the alerts")
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("DD-API-KEY", ddp.apiKey)
	req.Header.Add("DD-APPLICATION-KEY", ddp.appKey)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error occured when making http request")
	}

	fmt.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))

	result := DDMonitorSearchResponse{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error occured when parsing search response")
		fmt.Println(err)
	}

	return result
}
