package main

import (
	"client/utils"
	"fmt"
	"os"
	"strconv"
)

const (
	udp = "UDP"
	tcp = "TCP"
)

func main() {

	protocolType := os.Args[1]
	var times []int64

	numRequests, err := strconv.Atoi(os.Args[2])
	if err != nil {
		numRequests = 1
		fmt.Println("Invalid value for numRequests")
	}

	if protocolType == udp {
		fmt.Println("Running UDP client")
		times = RunUDPClient(numRequests)
	} else if protocolType == tcp {
		fmt.Println("Running TCP client")
		times = RunTCPClient(numRequests)
	} else {
		fmt.Println("Please select a valid protocol")
		return
	}

	mean := utils.GetMean(times)

	fmt.Println(mean)
}
