package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"resend/connect"
	"time"
)

type OrderStatus struct {
	Oid    int    `json:"oid"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func resend(ch *amqp.Channel, order *OrderStatus) error {
	//使用RabbitMQ发送订单到派送系统
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//装载消息
	body := []byte(order.Msg)
	//发送消息
	confirm, err := ch.PublishWithDeferredConfirmWithContext(ctx,
		//交换机
		"orderExchange",
		//路由key
		"",
		false,
		false,
		amqp.Publishing{ //设置消息体
			//MIME类型
			ContentType: "text/plain",
			//消息体
			Body: body,
		})
	if err != nil {
		return err
	}
	//接收确认
	if confirm.Wait() {
		return nil
	} else {
		return fmt.Errorf("send fail")
	}
}

func main() {
	//-----------------------环境初始化--------------------//
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
	//通道开启确认机制
	err = ch.Confirm(false)
	if err != nil {
		log.Println(err)
		return
	}
	//设置重发休眠时间
	sleepTime := time.Duration(1)
	//消息重发
	var resendList []OrderStatus
	for {
		var err error
		time.Sleep(sleepTime * time.Second)
		err = mysqlConn.Table("order_status").Where("status<0").Select("oid", "status").Find(&resendList).Error
		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}
		if len(resendList) == 0 {
			log.Println("No abnormal messages")
			time.Sleep(5 * time.Second)
			continue
		}
		for _, order := range resendList {
			err = resend(ch, &order)
			//发送成功则更新状态
			if err == nil {
				mysqlConn.Table("order_status").Where("oid=?", order.Oid).Update("status", 1)
			}
		}
	}
}
