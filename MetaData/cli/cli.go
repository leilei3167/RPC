package main

import (
	"context"
	"flag"
	"log"

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
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &example.Args{
		A: 10,
		B: 20,
	}
	reply := &example.Reply{}
	//将要发送的元数据加入上下文
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"aaa": "来自客户端"})
	//再构建一层,创建一个空map便于服务端写入
	ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
	//调用
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	//读取经服务端写过的元数据
	log.Printf("received meta: %+v", ctx.Value(share.ResMetaDataKey))

}
