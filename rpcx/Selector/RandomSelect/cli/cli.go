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

var (
	addr1 = flag.String("addr1", "tcp@localhost:9090", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:9091", "server address")
)

func main() {
	flag.Parse()

	//指定服务发现
	d, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1, Value: ""}, {Key: *addr2, Value: ""}})

	//创建客户端
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	//创建参数
	args := &example.Args{
		A: 10,
		B: 9,
	}

	//创建十个请求
	for i := 0; i < 10; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("结果%v为:%v\n", i, reply.C)
		time.Sleep(time.Second)
	}

}
