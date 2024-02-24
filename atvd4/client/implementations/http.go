package implementations

import (
	"atvd4/client/utils"
	"atvd4/common"
	"fmt"
	"net/rpc"
	"time"
)

func HttpRpcClient(board []int, iterations int, testing bool, boardNumber int) {
	var reply common.Reply

	client, err := rpc.DialHTTP("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	times := make([]float64, 0)

	for i := 0; i < iterations; i++ {

		startTime := time.Now()

		args := common.Request{
			Board: board,
		}

		err := client.Call("SudokuSolver.Run", args, &reply)

		if err != nil {
			panic(err)
		}

		duration := time.Since(startTime)

		times = append(times, float64(duration.Microseconds()))

		if testing {
			common.PrintBoard(reply.R, 9)
		}
	}

	utils.WriteDataInFile(times, "http", boardNumber)

	closeArgs := true
	var closeReply common.CloseReply

	err = client.Call("SudokuSolver.Close", closeArgs, &closeReply)

	if err != nil {
		fmt.Println("Ending application")
	}
}
