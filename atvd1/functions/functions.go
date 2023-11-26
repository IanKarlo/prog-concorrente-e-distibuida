package functions

import "math"

func ValidateSudoku(table [][]int, row int, col int, num int, size int) bool {
	return ValidateRow(table, row, num, size) && ValidateColumn(table, col, num, size) && ValidateBox(table, row, col, num, size)
}

func ValidateRow(table [][]int, row int, num int, size int) bool {
	for i := 0; i < size; i++ {
		if table[row][i] == num {
			return false
		}
	}
	return true
}

func ValidateColumn(table [][]int, col int, num int, size int) bool {
	for i := 0; i < size; i++ {
		if table[i][col] == num {
			return false
		}
	}
	return true
}

func ValidateBox(table [][]int, row int, column int, num int, size int) bool {
	boxes_number_f64 := math.Sqrt(float64(size))
	boxes_number_i32 := int(boxes_number_f64)
	box_row := int(math.Floor(float64(row) / boxes_number_f64))
	box_column := int(math.Floor(float64(column) / boxes_number_f64))

	box_row_start_index := box_row * boxes_number_i32
	box_column_start_index := box_column * boxes_number_i32

	for i := 0; i < boxes_number_i32; i++ {
		for j := 0; j < boxes_number_i32; j++ {
			if table[box_row_start_index+i][box_column_start_index+j] == num {
				return false
			}
		}
	}
	return true
}
