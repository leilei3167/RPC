package main

import (
	"context"
	"flag"
	"log"
	"sync"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8972", "server addr")

func main() {

	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	//创建链接池
	pool := client.NewXClientPool(10, "Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer pool.Close()
	args := &example.Args{
		A: 111,
		B: 1221,
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			reply := &example.Reply{}
			//需要从pool中取xclient出来,不需要放回
			xclient := pool.Get()
			err := xclient.Call(context.Background(), "Mul", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}

			log.Printf("%d * %d = %d", args.A, args.B, reply.C)
			wg.Done()
		}()

	}
	wg.Wait()

}
