package main

import (
	"context"
	"fmt"
	pb "gRPC/server/proto"
	"google.golang.org/grpc"
	"net"
)
type server struct {
	pb.UnimplementedHelloServer
}
func (s *server) SayHello(ctx context.Context,req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{MSG: "hello "+req.Name}, nil
}

func main(){
	//获取监听对象
	listener,err:=net.Listen("tcp","0.0.0.0:9090")
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//获取grpc服务对象
	grpcServer:=grpc.NewServer()
	//注册服务
	pb.RegisterHelloServer(grpcServer,&server{})
	//启动服务
	err=grpcServer.Serve(listener)
	if err!=nil{
		fmt.Println(err)
		return
	}
}