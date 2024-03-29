package implementations

import (
	"atvd4/server/impl"
	"net"
	"net/rpc"
)

func BuildTCPServer() {
	sudokuSolver := new(impl.SudokuSolver)

	server := rpc.NewServer()

	server.RegisterName("SudokuSolver", sudokuSolver)

	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic("Não foi posível ativar o listener")
	}

	server.Accept(listener)
}
