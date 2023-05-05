package main

import (
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
	//设置Qos
	err = ch.Qos(1, 0, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取消费者对象
	msgs, err := ch.Consume(
		"worker_queue",
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for data := range msgs {
			time.Sleep(time.Millisecond * 500)
			data.Ack(false)
			log.Printf("Received MSG: %s", string(data.Body))
		}
	}()
	fmt.Println("Waiting for message")
	var forever chan struct{}
	<-forever
}
