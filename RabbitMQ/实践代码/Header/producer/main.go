package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

var (
	IP       = "127.0.0.1"
	Port     = "5672"
	Account  = "admin"
	Password = "admin"
)

func main() {
	dataSource := fmt.Sprintf("amqp://%s:%s@%s:%s/", Account, Password, IP, Port)
	//获取连接
	conn, err := amqp.Dial(dataSource)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//获取通道对象
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()
	//设置超时关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//设置消息内容
	body := "hello world"
	//设置headers参数
	headers := make(map[string]interface{}, 1)
	headers["id"] = "x"
	//发送消息
	err = ch.PublishWithContext(ctx,
		//交换机
		"headX",
		//路由key
		"",
		false,
		false,
		amqp.Publishing{ //设置消息体
			//设置header
			Headers: headers,
			//MIME类型
			ContentType: "text/plain",
			//消息体
			Body: []byte(body),
		})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf(" SentMSG:%s\n", body)
}
