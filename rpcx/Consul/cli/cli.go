package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8080", "server addr")
var consuladdr = flag.String("consuladdr", "localhost:8500", "server addr")

func main() {
	flag.Parse()
	//consul服务发现模式
	d, _ := client.NewConsulDiscovery("/rpcx_test", "Arith", []string{*consuladdr}, nil)

	//创建客户端
	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	//调用

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &example.Reply{}

		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("调用成功:", reply.C)

		time.Sleep(time.Second * 2)

	}

}
