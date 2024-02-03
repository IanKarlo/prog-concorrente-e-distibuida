package utils

func GetMean(values []int64) float64 {
	total := 0.0

	for _, value := range values {
		total += float64(value)
	}

	if len(values) == 0 {
		return 0.0 // Evita a divisão por zero
	}

	return total / float64(len(values))
}
