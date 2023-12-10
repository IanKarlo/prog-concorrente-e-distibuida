package main

import (
	"atvd2/utils"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	finished bool
	s        int = 1
	r        int = 1
	mutex    sync.Mutex
)

var times []int64

const FULL_BIT_MASK = 511

func getSectorByRowAndColumn(row, col, size int) int {

	litteSquare := 3 //int(math.Sqrt(float64(size))) //3 depois muda pra melhorar o desepenho, qlqr coisa conta aqui irmão

	return col/litteSquare + litteSquare*(row/litteSquare)
}

func ValidateIfCanPutValue(row, col, value, size int, lines, columns, sectors []int) bool {
	sector := getSectorByRowAndColumn(row, col, size)

	return ((lines[row] | columns[col] | sectors[sector]) & value) == 0
}

func SolveRecursive(row, col int, board [][]int, size int, lines, columns, sectors []int) bool {

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
		return SolveRecursive(row, col, board, size, lines, columns, sectors)
	}

	for v := 1; v <= size; v++ {
		binValue := int(math.Pow(2, float64(v)-1.0))
		if ValidateIfCanPutValue(row, col, binValue, size, lines, columns, sectors) {
			board[row][col] = binValue
			setVectorCell(row, col, binValue, size, lines, columns, sectors)
			if SolveRecursive(row, col, board, size, lines, columns, sectors) {
				return true
			}
			removeVectorCell(row, col, binValue, size, lines, columns, sectors)
		}
	}

	board[row][col] = 0
	return false
}

func buildVectors(board [][]int, size int, lines, columns, sectors []int) {

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] != 0 {
				setVectorCell(i, j, board[i][j], size, lines, columns, sectors)
			}
		}
	}
}

func setVectorCell(row, col, binValue, size int, lines, columns, sectors []int) {
	sector := getSectorByRowAndColumn(row, col, size)

	lines[row] |= binValue
	columns[col] |= binValue
	sectors[sector] |= binValue
}

func removeVectorCell(row, col, binValue, size int, lines, columns, sectors []int) {
	maskedValue := FULL_BIT_MASK ^ binValue
	sector := getSectorByRowAndColumn(row, col, size)

	lines[row] &= maskedValue
	columns[col] &= maskedValue
	sectors[sector] &= maskedValue
}

func main() {

	numTries := 10000

	for i := 0; i < numTries; i++ {
		Run()
	}

	fmt.Println(times)
	fmt.Printf("%vµs\n", utils.GetMean(times))

}

func Run() {
	numRoutines := 5
	//fmt.Printf("Using %d goroutines\n", numRoutines)

	finished = false
	s = 1
	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go Solve(&wg, i)
	}

	wg.Wait()
}

func Solve(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	size := 9
	lines := make([]int, size)
	columns := make([]int, size)
	sectors := make([]int, size)

	// board := setUpBoard(size, []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0})
	board := setUpBoard(size, []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1})
	// board := setUpBoard(size, []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6})

	buildVectors(board, size, lines, columns, sectors)
	// printBoard(cellBoard, 0, size)
	// printBoardRow(cellBoard, 0, size)

	i, j, _ := getRandomValues(size)
	// printInitial(i, j, value, id)

	start := time.Now()

	if SolveRecursive(i, j, board, size, lines, columns, sectors) {
		// printBoard(board, id, size)
		end := time.Now()
		printTime(start, end)
		if !finished {
			finished = true
		}
	} else {
		fmt.Println("Deu ruimm")
	}
}

func setUpBoard(size int, puzzle []int) [][]int {
	newBoard := make([][]int, size)
	for i := range newBoard {
		newBoard[i] = make([]int, size)
	}

	n := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newBoard[i][j] = int(math.Pow(2, float64(puzzle[n])-1.0))
			n++
		}
	}

	return newBoard
}

func printTime(start, end time.Time) {
	mutex.Lock()
	defer mutex.Unlock()

	if s == 1 {
		s = 0

		times = append(times, end.Sub(start).Microseconds())
		//fmt.Printf("%v\n", end.Sub(start))
	}
}

// func printInitial(i, j, value int, id int) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	// fmt.Printf("Random row: %d, Random Col: %d, Random value: %d, Goroutine ID: %d\n", i, j, value, id)
// }

func getRandomValues(size int) (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(size - 1)
	j := rand.Intn(size - 1)
	value := rand.Intn(size-1) + 1
	return i, j, value
}

func printBoard(boardToPrint [][]int, id int, size int) {
	mutex.Lock()
	defer mutex.Unlock()

	base := int(math.Sqrt(float64(size)))
	separatorString := ""

	for i := 0; i < size; i++ {
		separatorString += "---"
	}

	if r == 1 {
		r = 1

		for i := 0; i < size; i++ {
			if i%base == 0 {
				fmt.Println(separatorString)
			}

			for j := 0; j < size; j++ {
				if j%base == 0 {
					fmt.Print("| ")
				}
				if boardToPrint[i][j] == 0 {
					fmt.Print("* ")
				} else {
					fmt.Printf("%.0f ", math.Log2(float64(boardToPrint[i][j]))+1)
				}
			}
			fmt.Println("|")
		}
		fmt.Println(separatorString)

		// fmt.Println("Goroutine that solved first:", id)
	}
}
