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
	addr1 = flag.String("addr1", "tcp@localhost:8972", "server address")
	addr2 = flag.String("addr2", "tcp@localhost:8973", "server address")
)

func main() {
	flag.Parse()

	d, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1, Value: ""},
		{Key: *addr2, Value: ""}})
	xclient := client.NewXClient("Arith", client.Failtry, client.ConsistentHash, d, client.DefaultOption)
	defer xclient.Close()

	args := &example.Args{
		A: 12,
		B: 10,
	}

	//会输出相同的结果
	for i := 0; i < 10; i++ {
		reply := &example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(time.Second)
	}

}
