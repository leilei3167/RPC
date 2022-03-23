package main

import (
	"context"
	"flag"
	"time"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	time.Sleep(time.Second)
	reply.C = args.A * args.B
	return nil
}
func main() {
	//设置server的读写超时时间
	s := server.NewServer(server.WithReadTimeout(time.Second*10),
		server.WithWriteTimeout(time.Second*10))
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)

}
