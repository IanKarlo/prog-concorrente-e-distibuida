package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"server/sudoku"
	"server/utils"
)

func StartUDPServer(maxIter int) {

	req := make([]byte, 1024)

	addr, err := net.ResolveUDPAddr("udp", "localhost:9091")

	if err != nil {
		fmt.Println("Error while resolving UDP Address")
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println("Error while creating UDP listener")
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < maxIter; i++ {
		_, addr, err := conn.ReadFromUDP(req)

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		limitIndex := utils.GetEndOfBuffer(req)

		var matrix []int
		err = json.Unmarshal([]byte(req[:limitIndex]), &matrix)

		res := sudoku.Run(i, matrix)

		jsonData, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		jsonData = append(jsonData, '\n')

		_, err = conn.WriteTo([]byte(jsonData), addr)

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}
	}

	err = conn.Close()

	if err != nil {
		fmt.Println("Error while accepting connection")
		os.Exit(1)
	}

}
