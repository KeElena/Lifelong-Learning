package main

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"producer/connect"
	"producer/init_env"
	"time"
)

type Order struct {
	Oid   int    `json:"oid"`
	Uid   int    `json:"uid"`
	Goods string `json:"goods"`
	Time  int64  `json:"time"`
}

type OrderStatus struct {
	Oid    int    `json:"oid"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
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
	//初始化队列和交换机
	err = init_env.InitEnv(ch)
	if err != nil {
		log.Println(err)
		return
	}
	//------------------用户发送订单---------------------------//
	//用户发送订单，订单存储到订单数据库
	userOrder := Order{Oid: int(time.Now().Unix() - 100000000), Uid: 1, Goods: "方便面", Time: time.Now().Unix()}
	err = mysqlConn.Create(&userOrder).Error
	if err != nil {
		log.Println(err)
		return
	}
	//------------------分布式事务之可靠发送-------------------//
	//获取订单内容
	msg, _ := json.Marshal(&userOrder)
	//订单状态存储到订单状态表
	err = mysqlConn.Table("order_status").Create(&OrderStatus{Oid: userOrder.Oid, Status: 0, Msg: string(msg)}).Error
	if err != nil {
		log.Println(err)
		return
	}
	//使用RabbitMQ发送订单到派送系统
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//装载消息
	body, _ := json.Marshal(&userOrder)
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
			Body:         body,
			DeliveryMode: 2,
		})
	if err != nil {
		log.Println(err)
		return
	}
	//接收确认
	if confirm.Wait() {
		//接收成功则更新状态
		mysqlConn.Table("order_status").Where("oid=?", userOrder.Oid).Update("status", 1)
	} else {
		mysqlConn.Table("order_status").Where("oid=?", userOrder.Oid).Update("status", -1)
	}
}
