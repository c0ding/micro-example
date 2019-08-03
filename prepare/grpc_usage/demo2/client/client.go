package main

import (
	"fmt"
	"google.golang.org/grpc"
)

const (
	post = ":8020"
)

func main() {
	var (
		conn *grpc.ClientConn
		err  error
	)
	if conn, err = grpc.Dial(post, grpc.WithInsecure()); err != nil {
		fmt.Println("连接服务器 错")
	}

	defer conn.Close()

}
