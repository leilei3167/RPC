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
	addr1 = flag.String("addr1", "tcp@localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "tcp@localhost:9981", "server2 address")
)

func main() {
	flag.Parse()
	d, err := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1},
		{Key: *addr2,Value: "group=test"}})
	if err != nil {
		log.Fatal(err)
	}
	//分组需要在客户端的option中进行
	op := client.DefaultOption
	op.Group = "test"

	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, op)
	defer xclient.Close()

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
		time.Sleep(time.Second * 3)

	}

}
