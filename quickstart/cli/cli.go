package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/leilei3167/RPC/quickstart/model"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", ":8972", "service address")
)

func main() {
	flag.Parse()
	//服务发现方式：点对点
	d, err := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	if err != nil {
		log.Fatal(err)
	}

	//创建client，服务路径为Arith，失败模式，随机选择服务器，服务发现，客户端默认选项
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &model.Args{A: 10, B: 20}
	var reply int

	//尝试异步调用
	arg := model.Args{A: 5, B: 4}

	call, err := xclient.Go(context.Background(), "Mul", arg, &reply, nil)
	if err != nil {
		log.Fatal(err)
	}
	//结果将会从channel中返回
	replycall := <-call.Done
	if replycall.Error != nil {
		log.Fatal(err)
	} else {
		fmt.Println("异步调用结果：", reply)
	}

	//调用服务器的Mul方法
	err = xclient.Call(context.Background(), "Mul", args, &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	fmt.Printf("%d * %d = %d\n", args.A, args.B, reply)

	//调用服务器的Div方法
	args = &model.Args{50, 20}
	var quo model.Quotient //会将结果写入这个结构体
	err = xclient.Call(context.Background(), "Div", args, &quo)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	fmt.Printf("%d * %d = %d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
