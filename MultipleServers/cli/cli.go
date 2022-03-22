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
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
)

func main() {
	flag.Parse()

	//服务发现要以MultipleServersDiscovery开始,将可选的地址以键值对形式放入
	d, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})

	//故障模式:选择另一个节点,选择器:使用roundrobin算法选择
	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
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
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}
