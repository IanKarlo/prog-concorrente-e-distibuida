package main

import (
	"atvd1/functions"
	"atvd1/utils"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	finished bool
	r        int = 1
	s        int = 1
	mutex    sync.Mutex
)

var times []int64

func main() {

	numTries := 1000

	for i := 0; i < numTries; i++ {
		run()
	}

	fmt.Println(times)
	fmt.Printf("%vµs\n", utils.GetMean(times))

}

func run() {
	numRoutines := 1
	//fmt.Printf("Using %d goroutines\n", numRoutines)

	finished = false
	s = 1
	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go solve(&wg, i)
	}

	wg.Wait()
}

func solve(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	size := 9

	board := setUpBoard(size, []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0})
	// board := setUpBoard(size, []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1})
	// board := setUpBoard(size, []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6})

	i, j, value := getRandomValues(size)
	// printInitial(i, j, value, id)

	start := time.Now()

	if solveRecursive(i, j, board, 0, value, size) {
		// printBoard(board, id, size)
		end := time.Now()
		printTime(start, end)
		if !finished {
			finished = true
		}
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
			newBoard[i][j] = puzzle[n]
			n++
		}
	}

	return newBoard
}

func solveRecursive(row, col int, board [][]int, xTimes, startV int, size int) bool {
	if xTimes == size*size {
		return true
	}

	isFull := true
	for p := 0; p < size; p++ {
		for q := 0; q < size; q++ {
			if board[p][q] == 0 {
				isFull = false
			}
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
		return solveRecursive(row, col, board, xTimes+1, startV, size)
	}

	for val := 1; val <= size; val++ {
		if startV++; startV == size+1 {
			startV = 1
		}

		if functions.ValidateSudoku(board, row, col, startV, size) {
			board[row][col] = startV
			if solveRecursive(row, col, board, xTimes+1, startV, size) {
				return true
			}
		}
	}

	board[row][col] = 0
	return false
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
		r = 0

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
					fmt.Printf("%d ", boardToPrint[i][j])
				}
			}
			fmt.Println("|")
		}
		fmt.Println(separatorString)

		// fmt.Println("Goroutine that solved first:", id)
	}
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

func printInitial(i, j, value int, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	// fmt.Printf("Random row: %d, Random Col: %d, Random value: %d, Goroutine ID: %d\n", i, j, value, id)
}

func getRandomValues(size int) (int, int, int) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(size - 1)
	j := rand.Intn(size - 1)
	value := rand.Intn(size-1) + 1
	return i, j, value
}
