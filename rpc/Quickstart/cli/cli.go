package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	//链接到服务端
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var r Result

	err := client.Call("HH.Add", 12, &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("同步值", r)

}
