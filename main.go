package main

import (
	"fmt"

	"github.com/vivasaayi/cloudrover/rovers"
	"github.com/vivasaayi/cloudrover/web"
)

func main() {
	fmt.Println("Starting Cloud Rover")

	fmt.Println("Starting Rovers")
	go rovers.StartRovers()

	fmt.Println("Starting Web Server")
	web.InitHttpServer()
}
