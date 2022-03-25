package main

import (
	"flag"
	"log"
	"time"

	example "github.com/leilei3167/RPC"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var addr = flag.String("addr", "localhost:8080", "server addr")
var consuladdr = flag.String("consuladdr", "localhost:8500", "server addr")
var basePath = flag.String("base", "/rpcx_test", "prefix path")

func main() {
	flag.Parse()
	s := server.NewServer()
	addConsulPlugin(s)
	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)
}

func addConsulPlugin(s *server.Server) {
	//配置插件
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consuladdr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	//链接
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	//添加到服务器
	s.Plugins.Add(r)

}
