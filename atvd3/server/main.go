package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	udp = "UDP"
	tcp = "TCP"
)

func main() {

	fmt.Println("Starting Server")

	protocolType := os.Args[1]

	numRequests, err := strconv.Atoi(os.Args[2])
	if err != nil {
		numRequests = 1
		fmt.Println("Invalid value for numRequests")
	}

	if protocolType == udp {
		fmt.Println("Running UDP server")
		StartUDPServer(numRequests)
	} else if protocolType == tcp {
		fmt.Println("Running TCP server")
		StartTCPServer(numRequests)
	} else {
		fmt.Println("Please select a valid protocol")
	}
}
