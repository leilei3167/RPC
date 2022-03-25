package main

import (
	"flag"

	"github.com/leilei3167/RPC/test/service"
	"github.com/smallnest/rpcx/server"
)

var addr1 = flag.String("addr1", "localhost:8080", "addr1")
var addr2 = flag.String("addr2", "localhost:8081", "addr2")
var addr3 = flag.String("addr3", "localhost:8082", "addr3")
var addr4 = flag.String("addr4", "localhost:8083", "addr4")

func main() {
	flag.Parse()
	//两个不同的地址注册相同的服务
	go createServer(*addr1, new(service.A))
	go createServer(*addr2, new(service.A))

	//另一个服务
	go createServer(*addr3, new(service.B))
	go createServer(*addr4, new(service.B))
	select {}

}

func createServer(addr string, recv interface{}) {
	s := server.NewServer()
	s.Register(recv, "") //注册A服务
	s.Serve("tcp", addr)

}
