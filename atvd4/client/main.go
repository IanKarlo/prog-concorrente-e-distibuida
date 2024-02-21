package main

import (
	"atvd4/common"
	"fmt"
	"net/rpc"
	"time"
)

func buildClient(board []int, iterations int) {
	var reply common.Reply

	client, err := rpc.Dial("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	mean := 0.0
	times := make([]float64, 0)

	for i := 0; i < iterations; i++ {

		startTime := time.Now()

		args := common.Request{
			Board: board,
		}

		err := client.Call("SudokuSolver.Run", args, &reply)

		if err != nil {
			fmt.Println(err)
		} else {
			// common.PrintBoard(reply.R, 9)
		}

		duration := time.Now().Sub(startTime)

		times = append(times, float64(duration.Microseconds()))

		mean += float64(duration.Microseconds())
	}

	fmt.Println(mean / float64(iterations))
	fmt.Println(times)
}

func main() {

	// board0 := []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0}
	board1 := []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1}
	// board2 := []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6}

	buildClient(board1, 10000)

}
