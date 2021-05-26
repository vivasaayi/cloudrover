package rovers

import (
	"fmt"

	"github.com/vivasaayi/cloudrover/rovers/datadog"
)

func StartRovers() {
	fmt.Println("Starting Alert Rover..")

	ar := datadog.GetDataDogAlertsRover()
	go ar.StartCollectingDataDogEvents()

	tmr := datadog.GetTriggeredMonitorsRover()
	go tmr.StartTriggeredMonitorsRover()

	fmt.Println("Rovers Started..")
}
