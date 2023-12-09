// package main

// import (
// 	"atvd2/utils"
// 	"fmt"
// 	"math"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// var (
// 	finished bool
// 	s        int = 1
// 	r        int = 1
// 	mutex    sync.Mutex
// )

// var times []int64

// type Cell struct {
// 	Value int
// 	Row   int
// 	Col   int
// 	Box   int
// 	Fixed bool
// }

// func validateIfCanPutValue(row, col, value int, board [][]Cell) bool {
// 	return ((board[row][col].Row | board[row][col].Col | board[row][col].Box) & value) == 0
// }

// func buildBoard(items [][]int, size int) [][]Cell {

// 	board := make([][]Cell, size)
// 	for temp := 0; temp < size; temp++ {
// 		board[temp] = make([]Cell, size)
// 	}

// 	for i := 0; i < size; i++ {
// 		for j := 0; j < size; j++ {
// 			var isFixed = false
// 			if items[i][j] != 0 {
// 				isFixed = true
// 			}
// 			setCell(i, j, items[i][j], isFixed, board, size)
// 		}
// 	}

// 	return board
// }

// func setCell(row, col, value int, fix bool, board [][]Cell, size int) {
// 	v := int(math.Pow(2, float64(value)-1.0))
// 	board[row][col].Value = v
// 	board[row][col].Fixed = fix
// 	sqrRootOfSize := int(math.Floor(math.Sqrt(float64(size))))
// 	for c := 0; c < size; c++ {
// 		board[row][c].Row |= v
// 		board[c][col].Col |= v
// 		board[(sqrRootOfSize*(row/sqrRootOfSize))+(c%sqrRootOfSize)][(sqrRootOfSize*(col/sqrRootOfSize))+(c/sqrRootOfSize)].Box |= v
// 	}
// }

// func newSolveRecursive(row, col int, board [][]Cell, size int) bool {

// 	isFull := true
// 	for p := 0; p < size; p++ {
// 		for q := 0; q < size; q++ {
// 			if board[p][q].Value == 0 {
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

// 	if board[row][col].Fixed {
// 		return newSolveRecursive(row, col, board, size)
// 	}

// 	for v := 1; v <= size; v++ {
// 		val := int(math.Pow(2, float64(v)-1.0))
// 		if validateIfCanPutValue(row, col, val, board) {
// 			board[row][col].Value = val
// 			if newSolveRecursive(row, col, board, size) {
// 				return true
// 			}
// 		}
// 	}

// 	board[row][col].Value = 0
// 	return false
// }

// func main() {

// 	numTries := 1

// 	for i := 0; i < numTries; i++ {
// 		newRun()
// 	}

// 	// fmt.Println(times)
// 	fmt.Printf("%vµs\n", utils.GetMean(times))

// }

// func newRun() {
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

// func newSolve(wg *sync.WaitGroup, id int) {
// 	defer wg.Done()

// 	size := 9

// 	// board := setUpBoard(size, []int{3, 0, 6, 5, 0, 8, 4, 0, 0, 5, 2, 0, 0, 0, 0, 0, 0, 0, 0, 8, 7, 0, 0, 0, 0, 3, 1, 0, 0, 3, 0, 1, 0, 0, 8, 0, 9, 0, 0, 8, 6, 3, 0, 0, 5, 0, 5, 0, 0, 9, 0, 6, 0, 0, 1, 3, 0, 0, 0, 0, 2, 5, 0, 0, 0, 0, 0, 0, 0, 0, 7, 4, 0, 0, 5, 2, 0, 6, 3, 0, 0})
// 	board := setUpBoard(size, []int{0, 0, 0, 0, 0, 5, 0, 6, 0, 0, 0, 5, 0, 6, 0, 0, 0, 0, 0, 7, 0, 0, 0, 0, 4, 0, 0, 0, 8, 0, 1, 0, 0, 0, 2, 0, 0, 3, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 5, 0, 2, 1, 0, 3, 0, 0, 6, 0, 0, 1, 0, 4, 0, 0, 0, 2, 0, 0, 0, 0, 7, 6, 0, 0, 3, 0, 0, 4, 5, 0, 1})
// 	// board := setUpBoard(size, []int{0, 9, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 4, 8, 0, 0, 5, 0, 6, 9, 2, 0, 0, 0, 0, 0, 9, 0, 0, 0, 0, 0, 0, 4, 2, 0, 0, 8, 0, 0, 0, 0, 8, 0, 7, 0, 0, 5, 0, 6, 1, 7, 0, 5, 9, 0, 4, 0, 4, 0, 0, 6, 0, 0, 5, 1, 2, 0, 0, 1, 0, 0, 0, 0, 6})

// 	cellBoard := buildBoard(board, size)
// 	// printBoard(cellBoard, 0, size)
// 	// printBoardRow(cellBoard, 0, size)

// 	i, j, _ := getRandomValues(size)
// 	// printInitial(i, j, value, id)

// 	start := time.Now()

// 	if newSolveRecursive(i, j, cellBoard, size) {
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

// func setUpBoard(size int, puzzle []int) [][]int {
// 	newBoard := make([][]int, size)
// 	for i := range newBoard {
// 		newBoard[i] = make([]int, size)
// 	}

// 	n := 0

// 	for i := 0; i < size; i++ {
// 		for j := 0; j < size; j++ {
// 			newBoard[i][j] = puzzle[n]
// 			n++
// 		}
// 	}

// 	return newBoard
// }

// func printTime(start, end time.Time) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	if s == 1 {
// 		s = 0

// 		times = append(times, end.Sub(start).Microseconds())
// 		fmt.Printf("%v\n", end.Sub(start))
// 	}
// }

// // func printInitial(i, j, value int, id int) {
// // 	mutex.Lock()
// // 	defer mutex.Unlock()

// // 	// fmt.Printf("Random row: %d, Random Col: %d, Random value: %d, Goroutine ID: %d\n", i, j, value, id)
// // }

// func getRandomValues(size int) (int, int, int) {
// 	rand.Seed(time.Now().UnixNano())
// 	i := rand.Intn(size - 1)
// 	j := rand.Intn(size - 1)
// 	value := rand.Intn(size-1) + 1
// 	return i, j, value
// }

// func printBoardRow(boardToPrint [][]Cell, id int, size int) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	base := int(math.Sqrt(float64(size)))
// 	separatorString := ""

// 	for i := 0; i < size; i++ {
// 		separatorString += "---"
// 	}

// 	if r == 1 {
// 		r = 0

// 		for i := 0; i < size; i++ {
// 			if i%base == 0 {
// 				fmt.Println(separatorString)
// 			}

// 			for j := 0; j < size; j++ {
// 				if j%base == 0 {
// 					fmt.Print("| ")
// 				}
// 				fmt.Printf("%09b ", boardToPrint[i][j].Row)
// 				// if boardToPrint[i][j].Value == 0 {
// 				// 	fmt.Print("* ")
// 				// } else {
// 				// 	fmt.Printf("%08b ", boardToPrint[i][j].Row)
// 				// }
// 			}
// 			fmt.Println("|")
// 		}
// 		fmt.Println(separatorString)

// 		// fmt.Println("Goroutine that solved first:", id)
// 	}
// }

// func printBoard(boardToPrint [][]Cell, id int, size int) {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	base := int(math.Sqrt(float64(size)))
// 	separatorString := ""

// 	for i := 0; i < size; i++ {
// 		separatorString += "---"
// 	}

// 	if r == 1 {
// 		r = 1

// 		for i := 0; i < size; i++ {
// 			if i%base == 0 {
// 				fmt.Println(separatorString)
// 			}

// 			for j := 0; j < size; j++ {
// 				if j%base == 0 {
// 					fmt.Print("| ")
// 				}
// 				if boardToPrint[i][j].Value == 0 {
// 					fmt.Print("* ")
// 				} else {
// 					fmt.Printf("%.0f ", math.Log2(float64(boardToPrint[i][j].Value))+1)
// 				}
// 			}
// 			fmt.Println("|")
// 		}
// 		fmt.Println(separatorString)

// 		// fmt.Println("Goroutine that solved first:", id)
// 	}
// }
