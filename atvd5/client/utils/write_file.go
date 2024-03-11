package utils

import (
	"fmt"
	"os"
	"strconv"
)

func WriteDataInFile(list []float64, protocol string, board int) error {

	fileName := fmt.Sprintf("%s_board%d_results.txt", protocol, board+1)

	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	mean := ArrayMean(list)
	median := ArrayMedian(list)

	fmt.Fprintf(file, "Média: %.2f;\n", mean)
	fmt.Fprintf(file, "Mediana: %.2f;\n", median)
	fmt.Fprintf(file, "Desvio padrão: %.2f;\n", median)

	arrayString := "["

	for el := range list {
		arrayString += strconv.FormatFloat(list[el], 'f', -1, 64) + ","
	}

	arrayStringLen := len(arrayString)

	arrayString = arrayString[:arrayStringLen-1] + "]"

	fmt.Fprintf(file, "Valores: %s;\n", arrayString)

	return nil

}
