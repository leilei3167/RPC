package main

import (
	"context"
	"errors"
	"flag"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server addr")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("Arith", new(example.Arith), "")
	//配置server的AuthFunc
	s.AuthFunc = auth

	s.Serve("tcp", *addr)

}

//AuthFunc的签名是固定的
func auth(ctx context.Context, req *protocol.Message, token string) error {
	if token == "bearer tGzv3JOkF0XG5Qx2TlKWIA" {

		return nil
	}
	return errors.New("无效的令牌!")

}
