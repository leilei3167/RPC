package main

import (
	"flag"

	gateway "github.com/rpcxio/rpcx-gateway"
	"github.com/rpcxio/rpcx-gateway/gin"
	"github.com/smallnest/rpcx/client"
)

var addr = flag.String("addr", "localhost:8972", "server addr")
var consuladdr = flag.String("consuladdr", "localhost:8500", "consul")

func main() {
	flag.Parse()
	d, _ := client.NewConsulDiscovery("/rpcx_test", "Arith", []string{*consuladdr}, nil)

	//创建新的网关,监听8091端口
	httpServer := gin.New(":8091")
	gw := gateway.NewGateway("/", httpServer, d, client.Failtry, client.RandomSelect, client.DefaultOption)

	gw.Serve()

}
