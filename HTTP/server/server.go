package main

import (
	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", "localhost:8080")

}
