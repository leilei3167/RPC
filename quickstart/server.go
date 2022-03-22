package main

import (
	"github.com/leilei3167/RPC/quickstart/model"

	"github.com/smallnest/rpcx/server"
)

func main() {
	//创建一个新服务
	s := server.NewServer()
	//向Arith路径注册Arith的所有方法
	s.RegisterName("Arith", new(model.Arith), "")
	s.Serve("tcp", ":8972")
	//以上是注册某个对象的方法到某个服务路径的形式,也可以直接注册函数
	//使用RegisterFunction

}
