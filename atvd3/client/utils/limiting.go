package utils

func GetEndOfBuffer(buff []byte) int {
	for i := 0; i < len(buff); i++ {
		if buff[i] == '\n' {
			return i
		}
	}
	return len(buff) - 1
}
