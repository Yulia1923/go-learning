package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-learning/day36/userpb"
)

type userServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(
	ctx context.Context,
	req *userpb.GetUserRequest,
) (*userpb.GetUserResponse, error) {
	fmt.Println("收到查询请求, 用户 ID: ", req.GetId())

	return &userpb.GetUserResponse{
		Id:   req.GetId(),
		Name: "Alice",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(
		grpcServer,
		&userServer{},
	)

	log.Println("gRPC 服务端正在监听: 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
