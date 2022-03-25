package main

import (
	"context"
	"flag"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B * 100
	return nil
}

func main() {
	flag.Parse()
	//分组为空和test
	go createServer1(*addr1, "group=nihao")
	go createServer2(*addr2, "group=test")

	select {}
}

func createServer1(addr, meta string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), meta)
	s.Serve("tcp", addr)

}

func createServer2(addr, meta string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), meta)
	s.Serve("tcp", addr)

}
