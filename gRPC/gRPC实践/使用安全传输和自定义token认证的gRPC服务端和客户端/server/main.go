package main

import (
	"context"
	"errors"
	"fmt"
	pb "gRPC/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
)

type server struct {
	pb.UnimplementedHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	//获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no Token")
	}
	//读取数据
	var appid, appkey string
	if v, ok := md["appid"]; ok {
		appid = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appkey = v[0]
	}
	//Token校验
	if appid != "demo" || appkey != "123456" {
		return nil, errors.New("token error")
	}
	return &pb.HelloResponse{MSG: "hello " + req.Name}, nil
}

func main() {
	//获取证书
	creds, err := credentials.NewServerTLSFromFile("../key/keqing.pem", "../key/keqing.key")
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取监听对象
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//获取grpc服务对象
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//注册服务
	pb.RegisterHelloServer(grpcServer, &server{})
	//启动服务
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
