package main

import (
	"os"
	"strconv"
)

func main() {

	numRequests, err := strconv.Atoi(os.Args[1])
	if err != nil {
		numRequests = 1
		panic("Invalid value for numRequests")
	}
	boardIndex, err := strconv.Atoi(os.Args[2])
	if err != nil {
		boardIndex = 1
		panic("Invalid value for boardIndex")
	}

	runClient(boardIndex, numRequests)

}
