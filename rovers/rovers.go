package rovers

import (
	"fmt"

	"github.com/vivasaayi/cloudrover/rovers/datadog"
)

func StartRovers() {
	fmt.Println("Starting Alert Rover..")
	ar := datadog.GetDataDogAlertsRover()
	ar.StartCollectingDataDogEvents()

	fmt.Println("Rovers Started..")
}
