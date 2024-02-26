package implementations

import (
	"atvd4/server/impl"
	"net"
	"net/http"
	"net/rpc"
)

func BuildHTTPServer() {
	sudokuSolver := new(impl.SudokuSolver)

	server := rpc.NewServer()

	server.RegisterName("SudokuSolver", sudokuSolver)

	server.HandleHTTP("/", "/debug")

	listener, err := net.Listen("tcp", ":5555")

	if err != nil {
		panic(err)
	}

	http.Serve(listener, nil)
}
