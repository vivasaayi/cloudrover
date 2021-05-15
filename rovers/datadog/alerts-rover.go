package datadog

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/vivasaayi/cloudrover/proxies"
	"github.com/vivasaayi/cloudrover/repositories"
)

type DataDogAlertsRover struct {
	ddProxy          *proxies.DataDogProxy
	previosEventTime time.Time
	visitedEvents    map[int64]bool
}

func GetDataDogAlertsRover() *DataDogAlertsRover {
	ar := DataDogAlertsRover{}

	ar.ddProxy = proxies.GetDataDogProxy()
	ar.previosEventTime = time.Now()
	ar.visitedEvents = map[int64]bool{}

	return &ar
}

func (ar *DataDogAlertsRover) StartCollectingDataDogEvents() {
	ticker := time.NewTicker(5 * time.Second)

	alerts := repositories.GetAllAlerts()
	for _, a := range alerts {
		ar.visitedEvents[a.Id] = true
	}

	for range ticker.C {
		ar.collectAndPublish()
	}
}

func (ar *DataDogAlertsRover) collectAndPublish() {
	fmt.Println("Publishing alert data")

	curTime := time.Now()

	startTime := ar.previosEventTime
	endTime := curTime

	ar.previosEventTime = curTime.Add(-8 * time.Hour)

	fmt.Println(startTime)
	fmt.Println(endTime)

	events := ar.ddProxy.GetEvents("datadog", startTime.Unix(), endTime.Unix(), "all")

	for _, event := range *events.Events {
		fmt.Println(*event.DateHappened)

		if _, ok := ar.visitedEvents[*event.Id]; ok {
			fmt.Printf("ID %d already visited\n", *event.Id)
			continue
		} else {
			ar.visitedEvents[*event.Id] = true
		}

		tags := ar.parseTags(&event)
		repositories.InsertDataDogAlert(&event, tags)
	}
}

func (ar *DataDogAlertsRover) parseTags(event *datadog.Event) string {
	data := map[string]string{}
	emptyJson := "{}"

	if event.Tags == nil {
		return emptyJson
	}

	for _, tag := range *event.Tags {
		st := strings.Split(tag, ":")

		key := st[0]
		value := ""

		if len(st) > 1 {
			value = st[1]
		}

		data[key] = value
	}

	jsonStr, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		return emptyJson
	}

	return string(jsonStr)
}
