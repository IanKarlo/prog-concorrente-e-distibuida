package sudoku

import (
	"math"
	"math/rand"
)

const FULL_BIT_MASK = 511

func SolveRecursive(row, col int, board [][]int, size int, lines, columns, sectors []int, signalChannel *chan int) bool {
	// tenta ler do canal bufferizado, se ler algo ele se encerra, se não ele continua
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

func Run(request int, board []int) [][]int {
	numRoutines := 5

	matrizChannel := make(chan [][]int)
	signalChannel := make(chan int, numRoutines)

	// fmt.Println("-----------", request)
	rand.Seed(42)
	for i := 0; i < numRoutines; i++ {
		go Solve(&matrizChannel, &signalChannel, board, i)
	}

	matrix := <-matrizChannel

	return matrix
}

func Solve(channel *chan [][]int, signalChannel *chan int, boardArray []int, id int) {

	size := 9
	lines := make([]int, size)
	columns := make([]int, size)
	sectors := make([]int, size)

	board := SetUpBoard(size, boardArray)

	BuildVectors(board, size, lines, columns, sectors)
	i, j, _ := getRandomValues(size)

	if SolveRecursive(i, j, board, size, lines, columns, sectors, signalChannel) {
		// fmt.Println("Sucesso :)", id)
		*signalChannel <- 1
		*channel <- board
		//manda informações de termino pro canal bufferizado
	} else {
		// fmt.Println("Encerrou sem sucesso", id)
	}
}

func getRandomValues(size int) (int, int, int) {
	// rand.Seed(time.Now().UnixNano())
	i := rand.Intn(size - 1)
	j := rand.Intn(size - 1)
	value := rand.Intn(size-1) + 1
	return i, j, value
}
