package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	op := client.DefaultOption
	op.IdleTimeout = time.Second * 10 //10秒最大空闲时长

	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, op)
	defer xclient.Close()

	//调用时令牌认证
	xclient.Auth("bearer 1tGzv3JOkF0XG5Qx2TlKWIA")

	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	//构建上下文
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatal("错误:", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
