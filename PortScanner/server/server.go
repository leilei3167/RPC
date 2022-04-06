package main

import (
	"context"
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

type Watcher struct {
}

type Task struct {
	TaskName string
}
type Res struct {
	Data string
}

func (w *Watcher) PutTask(ctx context.Context, arg *Task, res *Res) error {
	res.Data = arg.TaskName + "任务完成!!!"
	return nil
}

var (
	addr       = flag.String("addr", "localhost:8080", "监听地址")
	consulAddr = flag.String("consul", "localhost:8500", "consul地址")
	base       = flag.String("base", "/test", "前缀")
)

func main() {
	s := server.NewServer()

	addconsul(s)
	s.RegisterName("Watcher", new(Watcher), "")

	s.Serve("tcp", *addr)

}

func addconsul(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *base,
		UpdateInterval: time.Second * 30,
		Metrics:        metrics.NewRegistry(),
	}
	err := r.Start()
	if err != nil {
		log.Fatal("连接consul出错!", err)
	}

	s.Plugins.Add(r)

}
