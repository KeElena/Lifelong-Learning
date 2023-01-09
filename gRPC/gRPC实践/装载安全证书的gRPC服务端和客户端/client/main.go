package main

import (
	"context"
	"fmt"
	pb "gRPC/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//读取证书
	creds, err := credentials.NewClientTLSFromFile("../key/keqing.pem", "*.fosukeqing.cn")
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取grpc连接对象，设置不使用安全协议
	conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//获取客户端对象
	client := pb.NewHelloClient(conn)
	//客户端调用服务端的方法发送数据并获取响应
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "zhangsan"})
	if err != nil {
		fmt.Println(err)
		return
	}
	//处理响应数据
	fmt.Println(resp.GetMSG())
}
