package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func RunTCPClient(numRequests int) []int64 {
	times := make([]int64, 0)

	r, err := net.ResolveTCPAddr("tcp", "localhost:9091")

	if err != nil {
		fmt.Println("Error while accepting connection")
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, r)

	if err != nil {
		fmt.Println("Error while accepting connection")
		os.Exit(1)
	}

	requests := 0

	for requests < numRequests {

		start := time.Now()

		req := "Qlqr coisa"

		_, err = fmt.Fprintf(conn, req+"\n")

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		res, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		var matrix [][]int
		err = json.Unmarshal([]byte(res), &matrix)
		if err != nil {
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
