package main

import (
	"fmt"
	"github.com/c0ding/micro-example/prepare/protobuf_usage/demo1"
	"github.com/golang/protobuf/proto"
)

func main() {
	var (
		test *demo1.Test
		bytes []byte
		err error
		newtest *demo1.Test
	)

	test = &demo1.Test{
		Name:"老张",
		Stature:188,
		Weight: []int64{112,114,115,116,119},
		Motto:"0.1 * 365",
	}

	// 转换成protobuf
	if bytes, err = proto.Marshal(test); err != nil {
		fmt.Println("转码失败",err)
	}

	fmt.Println("转换成protobuf格式：",bytes)


	newtest = new(demo1.Test)
	if err = proto.Unmarshal(bytes, newtest); err != nil {
		fmt.Println("转码失败",err)
	}

	fmt.Println("protobuf格式 转换成 结构体：",newtest.String())
	fmt.Println("protobuf格式 转换成 结构体：",newtest)
	fmt.Println("protobuf格式 转换成 结构体里的字段：",newtest.GetName())
	fmt.Println("protobuf格式 转换成 结构体的字段：",newtest.Name)
	
}