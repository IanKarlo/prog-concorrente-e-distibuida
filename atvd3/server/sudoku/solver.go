package sudoku

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const FULL_BIT_MASK = 511

func SolveRecursive(row, col int, board [][]int, size int, lines, columns, sectors []int, signalChannel *chan int) bool {
	//tenta ler do canal bufferizado, se ler algo ele se encerra, se não ele continua
	if len(*signalChannel) != 0 {
		return false
	}

	isFull := true
	for p := 0; p < size; p++ {
		if lines[p] != FULL_BIT_MASK || columns[p] != FULL_BIT_MASK || sectors[p] != FULL_BIT_MASK {
			isFull = false
			break
		}
	}
	if isFull {
		return true
	}

	if col++; col == size {
		col = 0
		if row++; row == size {
			row = 0
		}
	}

	if board[row][col] != 0 {
		return SolveRecursive(row, col, board, size, lines, columns, sectors, signalChannel)
	}

	for v := 1; v <= size; v++ {
		binValue := int(math.Pow(2, float64(v)-1.0))
		if ValidateIfCanPutValue(row, col, binValue, size, lines, columns, sectors) {
			board[row][col] = binValue
			SetVectorCell(row, col, binValue, size, lines, columns, sectors)
			if SolveRecursive(row, col, board, size, lines, columns, sectors, signalChannel) {
				return true
			}
			RemoveVectorCell(row, col, binValue, size, lines, columns, sectors)
		}
	}

	board[row][col] = 0
	return false
}

func Run() [][]int {
	numRoutines := 5

	matrizChannel := make(chan [][]int)
	signalChannel := make(chan int, numRoutines)

	fmt.Println("-----------")
	for i := 0; i < numRoutines; i++ {
		go Solve(&matrizChannel, &signalChannel, i)
	}

	matrix := <-matrizChannel

	return matrix
}

func Solve(channel *chan [][]int, signalChannel *chan int, id int) {

	size := 9
	lines := make([]int, size)
	columns := make([]int, size)
	sectors := make([]int, size)

	// board := setUpBoard(size, []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0})
	board := SetUpBoard(size, []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1})
	// board := setUpBoard(size, []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6})

	BuildVectors(board, size, lines, columns, sectors)

	i, j, _ := getRandomValues(size)

	if SolveRecursive(i, j, board, size, lines, columns, sectors, signalChannel) {
		fmt.Println("Deu bom", id)
		*signalChannel <- 1
		*channel <- board
		//manda informações de termino pro canal bufferizado
	} else {
		fmt.Println("Deu ruimm", id)
	}
}

func getRandomValues(size int) (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(size - 1)
	j := rand.Intn(size - 1)
	value := rand.Intn(size-1) + 1
	return i, j, value
}
