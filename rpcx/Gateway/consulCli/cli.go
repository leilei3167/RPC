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

var consuladdr = flag.String("consuladdr", "localhost:8500", "consul")

func main() {
	flag.Parse()
	d, _ := client.NewConsulDiscovery("/rpcx_test", "Arith", []string{*consuladdr}, nil)

	args := &example.Args{
		A: 10,
		B: 20,
	}

	re := &example.Reply{}

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	err := xclient.Call(context.Background(), "Mul", args, re)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(re)
	time.Sleep(time.Second * 14)
}
