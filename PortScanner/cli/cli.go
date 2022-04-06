package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"log"
	"sync"
)

var (
	addr       = flag.String("addr", "localhost:8080", "监听地址")
	consulAddr = flag.String("consul", "localhost:8500", "consul地址")
	base       = flag.String("base", "/test", "前缀")
)

type Task struct {
	TaskName string
}
type Res struct {
	Data string
}

func main() {
	d, _ := client.NewConsulDiscovery(*base, "Watcher", []string{*consulAddr}, nil)
	client := client.NewXClient("Watcher", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	wg := sync.WaitGroup{}
	for i := 0; i < 150; i++ {
		x := i
		wg.Add(1)
		go func() {
			t1 := &Task{TaskName: string(x)}
			r1 := &Res{}

			err := client.Call(context.Background(), "PutTask", t1, r1)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}
			fmt.Println(r1.Data)

			wg.Done()
		}()

	}
	wg.Wait()

}
