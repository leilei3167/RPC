package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"time"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

var clientConn net.Conn
var connected = false

type Arith int

//net.Conn对象可以在客户端调用服务时获取,从ctx.Value(server.RemoteConnContextKey)获取
func (t *Arith) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	//获取conn
	clientConn = ctx.Value(server.RemoteConnContextKey).(net.Conn)
	reply.C = args.A * args.B
	connected = true
	return nil
}

func main() {
	flag.Parse()

	//	ln, _ := net.Listen("tcp", ":9981")
	//开启监听
	//	go http.Serve(ln, nil)
	//注册服务
	s := server.NewServer()
	//s.RegisterName("Arith", new(example.Arith), "")
	s.Register(new(Arith), "")
	go s.Serve("tcp", *addr)

	//如果未链接
	for !connected {
		time.Sleep(time.Second)

	}

	fmt.Printf("开始发送消息至%s \n", clientConn.RemoteAddr().String())

	for {
		if clientConn != nil {
			//s.SendMessage来发送消息给客户端
			err := s.SendMessage(clientConn, "test service path", "tets servic method", nil, []byte("abdcdcdcd"))
			if err != nil {
				fmt.Println("发送信息出错:", err)
				clientConn = nil
			}
		}
		//间隔1秒发送一次消息
		time.Sleep(time.Second)
	}

}
