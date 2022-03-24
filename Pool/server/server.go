package main

import (
	"context"
	"flag"
	"time"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith struct{}

// the second parameter is not a pointer
func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()

	//测试个请求率控制
	ratelimter := serverplugin.NewReqRateLimitingPlugin(time.Second, 10, true)
	s.Plugins.Add(ratelimter)
	//链接数控制
/* 	ratelimter2 := serverplugin.NewRateLimitingPlugin(time.Second, 1)
	s.Plugins.Add(ratelimter2) */
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
