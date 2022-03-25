package main

import (
	"context"
	"flag"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

//额外注册一个新服务
type Echo struct{}

func (t *Echo) Say(ctx context.Context, args string, reply *string) error {
	*reply = args
	return nil

}

func main() {

	flag.Parse()
	s := server.NewServer()
	s.Register(new(example.Arith), "")
	s.Register(new(Echo), "")
	s.Serve("tcp", "localhost:8080")

}
