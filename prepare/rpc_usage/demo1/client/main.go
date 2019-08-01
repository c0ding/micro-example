package main

import (
	"log"
	"net/rpc"
)

func main() {
	var (
		client *rpc.Client
		err    error
		args   int //上行参数
		result int //下行参数
	)
	if client, err = rpc.DialHTTP("tcp", ":8001"); err != nil {
		goto ERR
	}

	args = 1239
	if err = client.Call("Panda.GetInfo", args, &result); err != nil {
		goto ERR
	}

	log.Println("调用成功，返回值：", result)

	return
ERR:

	log.Fatal("c 失败", err.Error())
}
