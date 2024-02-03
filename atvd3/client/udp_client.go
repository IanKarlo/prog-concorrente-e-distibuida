package main

import (
	"client/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func RunUDPClient() {

	resBuffer := make([]byte, 1024)

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

	req := []byte("Qlqr coisa")

	_, err = conn.Write(req)

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
		return
	}

	utils.PrintBoard(matrix, 9)
}
