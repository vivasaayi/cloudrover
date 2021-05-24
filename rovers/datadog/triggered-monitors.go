package datadog

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/vivasaayi/cloudrover/proxies"
	"github.com/vivasaayi/cloudrover/repositories"
	"github.com/vivasaayi/cloudrover/utililties"
)

type TriggeredMonitorsRover struct {
	ddProxy          *proxies.DataDogProxy
	previosEventTime time.Time
	visitedEvents    map[int64]bool
}

func GetTriggeredMonitorsRover() *TriggeredMonitorsRover {
	tmr := TriggeredMonitorsRover{}

	tmr.ddProxy = proxies.GetDataDogProxy()
	tmr.previosEventTime = time.Now()
	tmr.visitedEvents = map[int64]bool{}

	return &tmr
}

func (tmr *TriggeredMonitorsRover) StartTriggeredMonitorsRover() {
	art := utililties.GetIntEnvVar("DD_TRIGGERED_MONITORS_SCHDULE", 60, false)
	ticker := time.NewTicker(time.Duration(art) * time.Second)

	for range ticker.C {
		tmr.ProduceMonitorsSummary()
	}
}

func (tmr *TriggeredMonitorsRover) ProduceMonitorsSummary() {
	fmt.Println("Producing Monitor Summary")

	monitors := tmr.ddProxy.SearchMonitors()
	checks := tmr.PrepareReport(monitors)

	jsonBody, err := json.MarshalIndent(checks, "", "  ")

	jsonStr := string(jsonBody)

	if err != nil {
		jsonStr = "{}"
	}

	repositories.InsertDataDogReport("triggered-monitors", 0, jsonStr)
}

func (tmr *TriggeredMonitorsRover) CheckMutedMonitors(mts []proxies.DDSearchResponseBoolSummary) CheckSummary {
	mutedMonitorsFound := false
	noOfMutedMonitors := 0

	for _, m := range mts {
		if m.Name {
			mutedMonitorsFound = true
			noOfMutedMonitors = m.Count
		}
	}

	chk := CheckSummary{
		Name: "No muted monitors.",
	}

	if mutedMonitorsFound {
		chk.Status = false
		chk.Count = noOfMutedMonitors
	}

	return chk
}

func (tmr *TriggeredMonitorsRover) PrepareReport(mts proxies.DDMonitorSearchResponse) []CheckSummary {
	checks := []CheckSummary{}

	mmchk := tmr.CheckMutedMonitors(mts.Counts.Muted)
	checks = append(checks, mmchk)

	fmt.Println("----------")
	fmt.Println(checks)
	fmt.Println("----------")

	return checks
}

func (tmr *TriggeredMonitorsRover) ProduceMonitorsSummaryI() {
	fmt.Println("Producing Monitor Summary")

	monitors := tmr.ddProxy.GetMonitors()

	for _, monitor := range monitors {
		fmt.Println(monitor)

		md := map[string]string{}

		md["Id"] = strconv.FormatInt(*monitor.Id, 8)
		md["Name"] = *monitor.Name
		md["Query"] = *monitor.Query

		// md["Priority"] = strconv.FormatInt(*monitor.Priority, 8)

		fmt.Printf("%+v\n", md)
	}
}
