package common

import (
	"fmt"
	"math/rand"
)

func HandleError(err error, msg string) {
	if err != nil {
		errMsg := fmt.Sprintf("%s!!: %s", msg, err)
		panic(errMsg)
	}
}

func GetConnectionSting() string {
	return fmt.Sprintf("amqp://%s:%s@localhost:5672/", amqpUser, amqpPassword)
}

func RandomId(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
