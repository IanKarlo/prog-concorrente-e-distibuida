package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"server/sudoku"
)

func StartTCPServer(maxIter int) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:9091")

	if err != nil {
		fmt.Println("Error while resolving TCP Address")
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		fmt.Println("Error while creating TCP listener")
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := listener.Accept()

	if err != nil {
		fmt.Println("Error while accepting connection")
		os.Exit(1)
	}

	for i := 0; i < maxIter; i++ {

		_, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		res := sudoku.Run()

		jsonData, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Error while accepting connection")
			os.Exit(1)
		}

		jsonData = append(jsonData, '\n')

		_, err = conn.Write([]byte(jsonData))

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
