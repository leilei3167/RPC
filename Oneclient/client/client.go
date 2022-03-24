package main

import (
	"context"
	"fmt"
	"log"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/client"
)

func main() {

	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8080", "")

	onecli := client.NewOneClient(client.Failtry, client.RoundRobin, d,
		client.DefaultOption)
	var echo string
	//一个oneclient可以调用不同的服务方法
	err := onecli.Call(context.Background(), "Echo", "Say", "dsadsadsadas", &echo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("结果1:", echo)

	args := &example.Args{
		A: 10,
		B: 20,
	}
	reply := &example.Reply{}
	err = onecli.Call(context.Background(), "Arith", "Mul", args, reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("结果2:", reply.C)

}
