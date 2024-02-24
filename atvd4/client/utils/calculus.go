package utils

import (
	"math"
	"sort"
)

func ArrayMean(array []float64) float64 {
	mean := 0.0

	for el := range array {
		mean += array[el]
	}

	return mean / float64(len(array))
}

func ArrayMedian(arr []float64) float64 {

	clonedList := make([]float64, len(arr))
	copy(clonedList, arr)

	sort.Float64s(clonedList)

	n := len(clonedList)
	if n == 0 {
		return 0
	}
	mid := n / 2
	if n%2 == 0 {
		return float64(clonedList[mid-1]+clonedList[mid]) / 2
	}
	return float64(clonedList[mid])
}

func ArrayStandardDeviation(arr []float64) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}

	mean := ArrayMean(arr)

	var squaredDiffs float64
	for _, num := range arr {
		squaredDiffs += math.Pow(num-mean, 2)
	}

	meanSquaredDiffs := squaredDiffs / float64(n)

	stdDev := math.Sqrt(meanSquaredDiffs)

	return stdDev
}
