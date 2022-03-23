package main

import (
	"flag"
	"net"
	"time"

	graphite "github.com/cyberdelia/go-metrics-graphite"
	example "github.com/leilei3167/RPC"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var addr = flag.String("addr", "localhost:8972", "server addr")

func main() {
	flag.Parse()

	//创建服务器
	s := server.NewServer()
	//配置插件

	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	s.Plugins.Add(p)
	startMetrics()

	s.Register(new(example.Arith), "")
	s.Serve("tcp", *addr)

}

func startMetrics() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Second)

	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	go graphite.Graphite(metrics.DefaultRegistry, time.Second, "rpcx.services.host.127_0_0_1", addr)

}
