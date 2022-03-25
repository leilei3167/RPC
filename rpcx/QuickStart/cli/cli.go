package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8080", "server addr")

func main() {
	flag.Parse()
	//创建服务发现的模式（Peer2peer等）
	d, err := client.NewPeer2PeerDiscovery(*addr, "")
	if err != nil {
		log.Fatal(err)
	}
	//构建新的客户端
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	//调用
	a := &example.Args{
		A: 10,
		B: 1000,
	}

	r := &example.Reply{}

	err = xclient.Call(context.Background(), "Mul", a, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功调用服务器注册在Arith的方法", r)

}
