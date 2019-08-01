package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Panda int

func (t *Panda) GetInfo(arge int, resp *int) error {

	log.Println("请求参数", arge)

	*resp = 1 + arge

	return nil
}

func main() {
	InitApiServer()

	//test()
}

func InitApiServer() {
	var (
		panda    *Panda
		listener net.Listener
		err      error
	)
	http.HandleFunc("/panda", handleList)
	panda = new(Panda)
	rpc.Register(panda)
	rpc.HandleHTTP()

	if listener, err = net.Listen("tcp", ":8001"); err != nil {
		goto ERR
	}

	http.Serve(listener, nil)

	return

ERR:
	log.Fatal("失败", err.Error())
}

func test() {
	var (
		panda *Panda

		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server

		err error
	)

	mux = http.NewServeMux()
	mux.HandleFunc("/panda", handleList)

	panda = new(Panda)
	rpc.Register(panda)
	rpc.HandleHTTP()

	if listener, err = net.Listen("tcp", ":8001"); err != nil {
		goto ERR
	}

	httpServer = &http.Server{
		Handler: mux,
	}
	httpServer.Serve(listener)

ERR:

	log.Fatal("失败", err.Error())

	//Register在默认服务中注册并公布 接收服务 panda对象 的方法
	// 其实就多了这两个步骤，其他和开启一个http 服务都一样的

}
func handleList(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "hello word  hello panda")
}
