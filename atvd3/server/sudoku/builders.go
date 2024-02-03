package sudoku

import (
	"math"
)

func SetUpBoard(size int, puzzle []int) [][]int {
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

func BuildVectors(board [][]int, size int, lines, columns, sectors []int) {

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] != 0 {
				SetVectorCell(i, j, board[i][j], size, lines, columns, sectors)
			}
		}
	}
}

func SetVectorCell(row, col, binValue, size int, lines, columns, sectors []int) {
	sector := GetSectorByRowAndColumn(row, col, size)

	lines[row] |= binValue
	columns[col] |= binValue
	sectors[sector] |= binValue
}

func RemoveVectorCell(row, col, binValue, size int, lines, columns, sectors []int) {
	maskedValue := FULL_BIT_MASK ^ binValue
	sector := GetSectorByRowAndColumn(row, col, size)

	lines[row] &= maskedValue
	columns[col] &= maskedValue
	sectors[sector] &= maskedValue
}
