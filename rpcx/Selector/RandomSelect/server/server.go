package main

import (
	"context"
	"flag"
	"time"

	example "github.com/leilei3167/RPC"

	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:9090", "server addr")
	addr2 = flag.String("addr2", "localhost:9091", "server addr")
)

//注册服务的三个条件:
//1.必须为可导出对象的方法 2.参数1必须是ctx 3.其余2个必须为可导出字段 4.最后一个必须为指针
type Arith struct{}

func (t *Arith) Mul(ctx context.Context, args *example.Args, rpl *example.Reply) error {
	rpl.C = args.A * args.B
	return nil

}

type Arith2 struct{}

func (t *Arith2) Mul(ctx context.Context, args *example.Args, rpl *example.Reply) error {
	rpl.C = args.A * args.B * 10000
	return nil

}

//注册
func main() {
	flag.Parse()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith), "")
		s.Serve("tcp", *addr1)
	}()

	go func() {
		s := server.NewServer()
		s.RegisterName("Arith", new(Arith2), "")
		s.Serve("tcp", *addr2)
	}()

	time.Sleep(time.Second * 1000)

}
