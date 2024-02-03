package main

import (
	"client/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func RunUDPClient(numRequests int, board []int) []int64 {
	resBuffer := make([]byte, 1024)
	times := make([]int64, 0)

	addr, err := net.ResolveUDPAddr("udp", "localhost:9091")

	if err != nil {
		fmt.Println("Error while accepting connection")
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)

	if err != nil {
		fmt.Println("Error while accepting connection")
		os.Exit(1)
	}

	requests := 0

	for requests < numRequests {

		start := time.Now()

		jsonData, err := json.Marshal(board)

		if err != nil {
			fmt.Println("Error while json.Marshal")
			os.Exit(1)
		}

		jsonData = append(jsonData, '\n')

		_, err = conn.Write(jsonData)

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		_, _, err = conn.ReadFromUDP(resBuffer)

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		limitIndex := utils.GetEndOfBuffer(resBuffer)

		var matrix [][]int
		err = json.Unmarshal([]byte(resBuffer[:limitIndex]), &matrix)
		if err != nil {
			fmt.Println(string(resBuffer))
			fmt.Println("Error:", err)
			panic(err)
		}

		// utils.PrintBoard(matrix, 9)

		end := time.Now()

		times = append(times, end.Sub(start).Microseconds())

		requests++
	}

	return times
}
