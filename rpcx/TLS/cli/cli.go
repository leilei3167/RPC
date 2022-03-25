package main

import (
	"context"
	"crypto/tls"
	"flag"
	"log"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	flag.Parse()

	//创建点对点服务发现,相当于"tcp@localhost:8972"
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	//配置选项
	option := client.DefaultOption
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	//配置default option的字段并创建XClient客户端
	option.TLSConfig = conf

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	//远程调用服务
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
}
