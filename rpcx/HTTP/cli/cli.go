package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	example "github.com/leilei3167/RPC"
	gateway "github.com/rpcxio/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
)

func main() {
	//msgpack编解码器
	cc := &codec.MsgpackCodec{}
	//参数
	args := &example.Args{
		A: 10,
		B: 1000,
	}

	//编码
	data, _ := cc.Encode(args)

	//创建req
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewReader(data))
	if err != nil {
		log.Fatal("创建请求错误:", req)

	}

	//必须设置header
	h := req.Header
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "Arith")
	h.Set(gateway.XServiceMethod, "Mul")

	//执行请求
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("发送请求失败!", err)
	}

	defer res.Body.Close()

	//获取结果
	replydata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := &example.Reply{}

	//解码
	err = cc.Decode(replydata, result)
	if err != nil {
		log.Fatal("解码错误", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, result.C)

}
