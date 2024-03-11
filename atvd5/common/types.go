package common

const RequestQueue = "request_queue"
const ResponseQueue = "response_queue"
const amqpUser = "user"
const amqpPassword = "password"

type Request struct {
	Board         []int
	ShouldTurnOff bool
}

type Reply struct {
	R [][]int
}

type CloseRequest bool

type CloseReply struct {
	Received bool
}
