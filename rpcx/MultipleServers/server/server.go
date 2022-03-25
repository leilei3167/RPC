package main

import (
	"flag"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

var (
	addr1 = flag.String("addr1", "localhost:8972", "server1 address")
	addr2 = flag.String("addr2", "localhost:9981", "server2 address")
)

func createServer(addr string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", addr)
}

func main() {
	//在两个不同的端口监听同一个服务
	go createServer(*addr1)
	go createServer(*addr2)

	select {}

}
