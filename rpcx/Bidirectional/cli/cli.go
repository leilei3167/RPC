package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	example "github.com/leilei3167/RPC"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	//需创建channel获取消息
	ch := make(chan *protocol.Message)

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	//需创建特殊的适用于双向数据交换的xclient
	xclient := client.NewBidirectionalXClient("Arith", client.Failtry,
		client.RandomSelect, d,
		client.DefaultOption,
		ch)
	defer xclient.Close()

	//调用服务
	args := &example.Args{
		A: 10,
		B: 20,
	}

	reply := &example.Reply{}
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	for msg := range ch {

		fmt.Printf("receive msg from server: %s\n", msg.Payload)

	}

}
