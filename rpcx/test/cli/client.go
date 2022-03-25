package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/leilei3167/RPC/test/service"
	"github.com/smallnest/rpcx/client"
)

var addr1 = flag.String("addr1", "localhost:8080", "addr1")
var addr2 = flag.String("addr2", "localhost:8081", "addr2")
var addr3 = flag.String("addr3", "localhost:8082", "addr3")
var addr4 = flag.String("addr4", "localhost:8083", "addr4")

func main() {
	flag.Parse()
	//不同的服务必须单独开xclient
	d1, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	d2, _ := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr3}, {Key: *addr4}})

	xclient1 := client.NewXClient("A", client.Failtry, client.RoundRobin, d1, client.DefaultOption)
	xclient2 := client.NewXClient("B", client.Failtry, client.RoundRobin, d2, client.DefaultOption)

	in := &service.In{
		I: "nihao",
	}

	for i := 0; i < 5; i++ {
		go func() {
			r := &service.R{}
			err := xclient1.Call(context.Background(), "Addhello", in, r)
			if err != nil {
				log.Printf("调用服务Addhello出错!\n", err)
			}
			fmt.Printf("执行addhello成功:%v\n", r.S)
			time.Sleep(time.Second)
		}()
		go func() {
			r := &service.R{}
			err := xclient1.Call(context.Background(), "Addgin", in, r)
			if err != nil {
				log.Printf("调用服务Addgin出错!\n", err)
			}
			fmt.Printf("执行addgin成功:%v\n", r.S)
			time.Sleep(time.Second)
		}()

		go func() {
			r := &service.R{}
			err := xclient2.Call(context.Background(), "Add1", in, r)
			if err != nil {
				log.Printf("调用服务Add1出错!\n", err)
			}
			fmt.Printf("执行add1成功:%v\n", r.S)
			time.Sleep(time.Second)
		}()
		go func() {
			r := &service.R{}
			err := xclient2.Call(context.Background(), "Add2", in, r)
			if err != nil {
				log.Printf("调用服务Add2n出错!\n", err)
			}
			fmt.Printf("执行add2成功:%v\n", r.S)
			time.Sleep(time.Second)
		}()

	}
	time.Sleep(time.Second * 4)

}
