package impl

import (
	"atvd5/common"
	"math/rand"
)

type SudokuSolver struct{}

func (s *SudokuSolver) Run(req common.Request) [][]int {

	board := req.Board

	numRoutines := 5

	matrizChannel := make(chan [][]int)
	signalChannel := make(chan int, numRoutines)

	// fmt.Println("-----------", request)
	rand.Seed(42)
	for i := 0; i < numRoutines; i++ {
		go Solve(&matrizChannel, &signalChannel, board, i)
	}

	matrix := <-matrizChannel

	// PrintBoard(matrix, 9)

	return matrix
}