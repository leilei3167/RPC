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
	addr = flag.String("addr", "localhost:8972", "server address")
	hb   = flag.Bool("hb", true, "enable heartbeat or not")
)

func main() {
	flag.Parse()
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	//配置客户端选项
	op := client.DefaultOption
	//开启心跳,周期性发送心跳信息到服务端并等待回应
	op.Heartbeat = *hb
	//心跳间隔1秒,最大心跳等待时间2秒
	op.HeartbeatInterval = time.Second
	op.MaxWaitForHeartbeat = 2 * time.Second

	op.IdleTimeout = 3 * time.Second //net.Conn最大的空闲时间
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, op)
	defer xclient.Close()

	args := &example.Args{
		A: 10,
		B: 20,
	}

	for i := 0; i < 10; i++ {
		reply := &example.Reply{}

		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)

		time.Sleep(10 * time.Second)

	}

}
