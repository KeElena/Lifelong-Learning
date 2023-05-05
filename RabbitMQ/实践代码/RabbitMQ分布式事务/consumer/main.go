package main

import (
	"consumer/connect"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Dispatch struct {
	Oid      int    `json:"oid"`
	Delivery string `json:"delivery"`
	Status   string `json:"status"`
}

func main() {
	//获取连接
	rabbitConn, err := connect.GetRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}
	mysqlConn, err := connect.GetMysql()
	if err != nil {
		log.Println(err)
		return
	}
	//获取通道对象
	ch, err := rabbitConn.Channel()
	if err != nil {
		log.Println(err)
		return
	}
	defer ch.Close()
	//获取消息
	msgs, err := ch.Consume(
		"orderQueue",
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println(err)
		return
	}
	//处理消息
	go func() {
		for order := range msgs {
			time.Sleep(time.Millisecond * 500)
			var deliver Dispatch
			_ = json.Unmarshal(order.Body, &deliver)
			deliver.Delivery = "刘强东"
			deliver.Status = "配送中"
			err := mysqlConn.Table("dispatch").Create(&deliver).Error
			if err != nil {
				//requeue设为true时重新分发，false时转移到死信队列
				_ = order.Nack(false, false)
				continue
			}
			_ = order.Ack(false)
			log.Println("成功处理一份订单")
		}
	}()
	fmt.Println("Waiting for message")
	var forever chan struct{}
	<-forever
}
