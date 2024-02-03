package main

import (
	"bufio"
	"client/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func RunTCPClient() {
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
		return
	}

	utils.PrintBoard(matrix, 9)
}
