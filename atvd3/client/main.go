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
	board0 := []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0}
	board1 := []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1}
	board2 := []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6}

	boards := [][]int{board0, board1, board2}

	protocolType := os.Args[1]
	var times []int64

	numRequests, err := strconv.Atoi(os.Args[2])
	if err != nil {
		numRequests = 1
		fmt.Println("Invalid value for numRequests")
	}
	boardIndex, err := strconv.Atoi(os.Args[3])
	if err != nil {
		boardIndex = 1
		fmt.Println("Invalid value for boardIndex")
	}

	if protocolType == udp {
		fmt.Println("Running UDP client")
		times = RunUDPClient(numRequests, boards[boardIndex])
	} else if protocolType == tcp {
		fmt.Println("Running TCP client")
		times = RunTCPClient(numRequests, boards[boardIndex])
	} else {
		fmt.Println("Please select a valid protocol")
		return
	}

	mean := utils.GetMean(times)
	fmt.Println(times)

	fmt.Println(mean)
}
