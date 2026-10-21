package main

import (
	"fmt"
	"go-learning/day35/userpb"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &userpb.User{
		Id:   1001,
		Name: "Alice",
		Age:  20,
	}

	data, err := proto.Marshal(user)
	if err != nil {
		panic(err)
	}

	decoded := &userpb.User{} //创建了一个新的、初始字段为空或零值的User对象
	err = proto.Unmarshal(data, decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("序列化后的字节长度:", len(data))
	fmt.Println("序列化是否成功:", err == nil)
	fmt.Println("还原后的ID: ", decoded.GetId())
	fmt.Println("还原后的Name:", decoded.GetName())
	fmt.Println("还原后的Age: ", decoded.GetAge())
}
