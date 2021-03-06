package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UManagerServiceServer ...
type UManagerServiceServer struct{
	Users []*User
}

// GetUser ...
func (server *UManagerServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (res *GetUserResponse, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			return &GetUserResponse{
				Name:    u.Name,
				Title:   u.Title,
				Company: u.Company,
			}, nil
		}
	}
	return nil, ErrUserNotFound
}

// SetTitle ...
func (server *UManagerServiceServer) SetTitle(ctx context.Context, req *SetTitleRequest) (res *Empty, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			u.Title = req.Title
			return &Empty{}, nil
		}
	}
	return &Empty{}, ErrUserNotFound
}

// Dispatch ...
func (server *UManagerServiceServer) Dispatch(ctx context.Context, req *DispatchRequest) (res *Empty, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			u.Company = req.Company
			return &Empty{}, nil
		}
	}
	return &Empty{}, ErrUserNotFound
}

// NewServer ...
func NewServer() {
	log.Println("server: 启动监听")
	lis, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		panic(err)
	}
	rpcServer := grpc.NewServer()
	log.Println("server: 注册服务")
	RegisterUserManagerServiceServer(rpcServer, uManagerServiceServer)
	reflection.Register(rpcServer)
	log.Println("server: 等待连接")
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
