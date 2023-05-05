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
	//声明业务交换机
	err = ch.ExchangeDeclare("task_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//声明死信交换机
	err = ch.ExchangeDeclare("dead_exchange", "direct", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//声明队列
	//声明过期队列
	ttl_args := make(map[string]interface{}, 2)
	//设置过期数据
	ttl_args["x-message-ttl"] = 5000
	//设置死信交换机
	ttl_args["x-dead-letter-exchange"] = "dead_exchange"
	//设置死信路由key，fanout模式不用设置
	ttl_args["x-dead-letter-routing-key"] = "dead"
	_, err = ch.QueueDeclare("ttl_queue", false, false, false, false, ttl_args)
	if err != nil {
		fmt.Println(err)
		return
	}
	//声明死信队列
	_, err = ch.QueueDeclare("dead_queue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//绑定队列
	//死信队列绑定到死信交换机
	err = ch.QueueBind("dead_queue", "dead", "dead_exchange", false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//过期队列绑定到业务交换机
	err = ch.QueueBind("ttl_queue", "", "task_exchange", false, nil)
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
			"task_exchange",
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
