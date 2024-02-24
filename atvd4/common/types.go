package common

type Request struct {
	Board []int
}

type Reply struct {
	R [][]int
}

type CloseRequest bool

type CloseReply struct {
	Received bool
}
