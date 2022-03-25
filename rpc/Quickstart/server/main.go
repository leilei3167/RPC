package main

import (
	"log"
	"net/http"
	"net/rpc"
)

/* 注册为RPC服务的几个必要条件:
func (t *T) MethodName(argType T1, replyType *T2) error
1.类型T必须是可导出的
2.服务名必须是可导出的
3.两个参数必须为可导出的
4.第二个参数必须是指针
5.返回值必须是error
*/

//创建一个服务
type Result struct {
	Num, Ans int
}
type Cal struct{}

func (c *Cal) Squ(num int, res *Result) error {
	res.Num = num
	res.Ans = num * num

	return nil

}
func (c *Cal) Add(num int, res *Result) error {
	res.Num = num
	res.Ans = num + num

	return nil

}

func main() {
	//注册服务
	rpc.RegisterName("HH", new(Cal))
	//注册一个用于处理RPC消息的http Handler
	rpc.HandleHTTP()

	if err := http.ListenAndServe(":1234", nil); err != nil {

		log.Fatal(err)
	}

}
