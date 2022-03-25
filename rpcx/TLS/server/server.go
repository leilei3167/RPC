package main

import (
	"crypto/tls"
	"flag"
	"log"

	example "github.com/leilei3167/RPC"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	//此处载入证书和私钥(可生成)
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Print(err)
		return
	}
	//将解析的加密文件加载到server配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	s := server.NewServer(server.WithTLSConfig(config))

	s.RegisterName("Arith", new(example.Arith), "")
	s.Serve("tcp", *addr)

}
