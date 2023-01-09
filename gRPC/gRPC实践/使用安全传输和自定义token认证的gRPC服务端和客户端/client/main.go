package main

import (
	"context"
	"fmt"
	pb "gRPC/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//复制接口
type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

//定义Token结构体
type ClientToken struct {
}

//设置发送的Token信息
func (c *ClientToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "demo",
		"appkey": "123456",
	}, nil
}

//设置是否开启证书安全传输认证
func (c *ClientToken) RequireTransportSecurity() bool {
	return true
}

func main() {
	//读取证书
	creds, err := credentials.NewClientTLSFromFile("../key/keqing.pem", "*.fosukeqing.cn")
	if err != nil {
		fmt.Println(err)
		return
	}
	//配置grpc客户端
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientToken)))
	//获取grpc连接对象，设置不使用安全协议
	conn, err := grpc.Dial("0.0.0.0:9090", opts...)
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
