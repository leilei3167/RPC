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

var (
	//定义一个服务器地址
	addr = flag.String("addr", "localhost:8972", "server address")
	//定义一个zookeeper集群地址
	zkAddr = flag.String("zkAddr", "localhost:2181", "zookeeper address")
	//
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	addRegistryPlugin(s) //配置插件

	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)

}

//配置插件的函数
func addRegistryPlugin(s *server.Server) {
	//配置Zookeeper插件
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,    //服务地址
		ZooKeeperServers: []string{*zkAddr}, //集群地址
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	//链接到集群
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	//将插件添加到server的Plugins字段
	s.Plugins.Add(r)

}
