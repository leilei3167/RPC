package main

import (
	"flag"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

//将地址定义为常量
var addr = flag.String("addr", "localhost:8080", "server addr")

func main() {
	flag.Parse()
	//生成服务器
	s := server.NewServer()
	//注册Arith对象的方法
	s.Register(new(example.Arith), "") //不指定服务名称，默认为对象名
	//s.RegisterName("Arith", new(example.Arith), "")

	s.Serve("tcp", *addr)

}
