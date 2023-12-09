// package main

// import (
// 	"atvd2/utils"
// 	"fmt"
// 	"math"
// 	"sync"
// 	"time"
// )

// func getSectorByRowAndColumn(row, col, size int) int {

// 	litteSquare := int(math.Sqrt(float64(size))) //3 depois muda pra melhorar o desepenho, qlqr coisa conta aqui irmão

// 	return col/litteSquare + litteSquare*(row/litteSquare)
// }

// func newValidateIfCanPutValue(row, col, value, size int, lines, columns, sectors []int) bool {
// 	sector := getSectorByRowAndColumn(row, col, size)

// 	return ((lines[row] | columns[col] | sectors[sector]) & value) == 0
// }

// func SolveRecursive(row, col int, board [][]int, size int, lines, columns, sectors []int) bool {

// 	isFull := true
// 	for p := 0; p < size; p++ {
// 		for q := 0; q < size; q++ {
// 			if board[p][q] == 0 {
// 				isFull = false
// 			}
// 		}
// 	}
// 	if isFull {
// 		return true
// 	}

// 	if col++; col == size {
// 		col = 0
// 		if row++; row == size {
// 			row = 0
// 		}
// 	}

// 	if board[row][col] != 0 {
// 		return SolveRecursive(row, col, board, size, lines, columns, sectors)
// 	}

// 	for v := 1; v <= size; v++ {
// 		binValue := int(math.Pow(2, float64(v)-1.0))
// 		if newValidateIfCanPutValue(row, col, binValue, size, lines, columns, sectors) {
// 			board[row][col] = binValue
// 			if SolveRecursive(row, col, board, size, lines, columns, sectors) {
// 				return true
// 			}
// 		}
// 	}

// 	board[row][col] = 0
// 	return false
// }

// func buildVectors(board [][]int, size int, lines, columns, sectors []int) {

// 	for i := 0; i < size; i++ {
// 		for j := 0; j < size; j++ {
// 			if (board[i][j] != 0) {
// 				setVectorCell(i, j, board[i][j], size, board, lines, columns, sectors)
// 			}
// 		}
// 	}
// }

// func setVectorCell(row, col, value, size int, board [][]int, lines, columns, sectors []int) {
// 	binValue := int(math.Pow(2, float64(value)-1.0))
// 	sector := getSectorByRowAndColumn(row, col, size)

// 	lines[row] |= binValue
// 	columns[col] |= binValue
// 	sectors[sector] |= binValue

// }

// func main() {

// 	numTries := 1

// 	for i := 0; i < numTries; i++ {
// 		Run()
// 	}

// 	// fmt.Println(times)
// 	fmt.Printf("%vµs\n", utils.GetMean(times))

// }

// func Run() {
// 	numRoutines := 1
// 	//fmt.Printf("Using %d goroutines\n", numRoutines)

// 	finished = false
// 	s = 1
// 	var wg sync.WaitGroup
// 	wg.Add(numRoutines)

// 	for i := 0; i < numRoutines; i++ {
// 		go newSolve(&wg, i)
// 	}

// 	wg.Wait()
// }

// func Solve(wg *sync.WaitGroup, id int) {
// 	defer wg.Done()

// 	size := 9
// 	lines := make([]int, size)
// 	columns := make([]int, size)
// 	sectors := make([]int, size)

// 	// board := setUpBoard(size, []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0})
// 	board := setUpBoard(size, []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1})
// 	// board := setUpBoard(size, []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6})

// 	buildVectors(board, size, lines, columns, sectors)
// 	// printBoard(cellBoard, 0, size)
// 	// printBoardRow(cellBoard, 0, size)

// 	i, j, _ := getRandomValues(size)
// 	// printInitial(i, j, value, id)

// 	start := time.Now()

// 	if SolveRecursive(i, j, board, size, lines, columns, sectors) {
// 		// printBoard(cellBoard, id, size)
// 		end := time.Now()
// 		printTime(start, end)
// 		if !finished {
// 			finished = true
// 		}
// 	} else {
// 		fmt.Println("Deu ruimm")
// 	}
// }
