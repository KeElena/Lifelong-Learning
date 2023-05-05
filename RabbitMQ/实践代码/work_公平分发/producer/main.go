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
	//声明交换机
	err = ch.ExchangeDeclare("work_fair", "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//声明队列
	_, err = ch.QueueDeclare("worker_queue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//队列绑定到交换机
	err = ch.QueueBind("worker_queue", "", "work_fair", false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置超时关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//发送消息
	n := 0
	for {
		//手动设置模拟处理时间
		if n == 40 {
			break
		}
		time.Sleep(100 * time.Millisecond)
		body := fmt.Sprintf("%d", n)
		//发送消息
		err = ch.PublishWithContext(ctx,
			//交换机
			"work_fair",
			//路由key
			"",
			false,
			false,
			amqp.Publishing{ //设置消息体
				//MIME类型
				ContentType: "text/plain",
				//消息体
				Body: []byte(body),
			})
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("SentMSG:%s\n", body)
		n++
	}
}
