package init_env

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitEnv(ch *amqp.Channel) error {
	var err error
	//声明死信交换机
	err = ch.ExchangeDeclare("deadExchange", "direct", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("deadExchange:" + err.Error())
	}
	//声明消息队列
	args := make(map[string]interface{})
	args["x-dead-letter-exchange"] = "deadExchange"
	args["x-dead-letter-routing-key"] = "order"
	_, err = ch.QueueDeclare("orderQueue", true, false, false, false, args)
	if err != nil {
		return fmt.Errorf("ordermessage:" + err.Error())
	}
	//声明死信队列
	_, err = ch.QueueDeclare("deadMessage", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("deadQueue:" + err.Error())
	}
	//死信队列绑定
	err = ch.QueueBind("deadMessage", "order", "deadExchange", false, nil)
	if err != nil {
		return fmt.Errorf("deadMessage QueueBind:" + err.Error())
	}
	//声明分发订单交换机
	err = ch.ExchangeDeclare("orderExchange", "direct", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("orderExchange:" + err.Error())
	}
	//订单队列绑定
	err = ch.QueueBind("orderQueue", "", "orderExchange", false, nil)
	if err != nil {
		return fmt.Errorf("orderQueue QueueBind:" + err.Error())
	}
	return nil
}
