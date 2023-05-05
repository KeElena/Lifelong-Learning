package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
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

	//获取消息
	msgs, err := ch.Consume(
		"two",
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
			log.Printf("Received MSG: %s", data.Body)
		}
	}()
	fmt.Println("Waiting for message")
	var forever chan struct{}
	<-forever
}
