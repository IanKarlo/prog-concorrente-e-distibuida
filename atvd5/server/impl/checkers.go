package impl

func GetSectorByRowAndColumn(row, col, size int) int {

	litteSquare := 3 //int(math.Sqrt(float64(size))) //3 depois muda pra melhorar o desepenho, qlqr coisa conta aqui irmão

	return col/litteSquare + litteSquare*(row/litteSquare)
}

func ValidateIfCanPutValue(row, col, value, size int, lines, columns, sectors []int) bool {
	sector := GetSectorByRowAndColumn(row, col, size)

	return ((lines[row] | columns[col] | sectors[sector]) & value) == 0
}
