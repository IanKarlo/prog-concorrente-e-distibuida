package main

import (
	"atvd4/server/implementations"
	"os"
)

func main() {
	protocolType := os.Args[1]

	if protocolType == "TCP" {
		implementations.BuildTCPServer()
	} else {
		implementations.BuildHTTPServer()
	}
}
