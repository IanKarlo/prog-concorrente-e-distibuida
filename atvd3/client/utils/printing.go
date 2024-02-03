package utils

import (
	"fmt"
	"math"
)

func PrintBoard(boardToPrint [][]int, size int) {

	base := int(math.Sqrt(float64(size)))
	separatorString := ""

	for i := 0; i < size; i++ {
		separatorString += "---"
	}

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
}
