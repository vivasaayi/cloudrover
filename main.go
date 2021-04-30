package main

import (
	"fmt"

	"github.com/vivasaayi/cloudrover/web"
)

func main() {
	fmt.Println("Starting Cloud Rover")
	web.InitHttpServer()
}
