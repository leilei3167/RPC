package main

import (
	"context"
	"flag"
	"fmt"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type Arith int

//构建服务
func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	//读取context中的元数据
	reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)

	fmt.Println("收到元数据:", reqMeta)

	resMeta["echo"] = "来自服务器的元数据"

	reply.C = args.A * args.B
	return nil

}

func main() {

	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}
