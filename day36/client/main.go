package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"go-learning/day36/userpb"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn) //创建生成代码提供的客户端

	req := &userpb.GetUserRequest{
		Id: 1001,
	}

	resp, err := client.GetUser(ctx, req) //发起真正的RPC调用
	if err != nil {
		if status.Code(err) == codes.NotFound {
			log.Println("用户不存在")
			return
		}

		log.Printf("RPC 调用失败：%v\n", err)
		return
	}

	fmt.Printf(
		"查询成功: 用户 ID = %d, 姓名 = %s\n",
		resp.GetId(),
		resp.GetName(),
	)
}
