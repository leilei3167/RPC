package main

import (
	"context"
	"flag"
	"log"
	"time"

	example "github.com/leilei3167/RPC"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}

	//构建客户端等待指定时长超时的context
	ctx, cancleFunc := context.WithTimeout(context.Background(), time.Second*1)

	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatal(err)
	}
	cancleFunc()

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
