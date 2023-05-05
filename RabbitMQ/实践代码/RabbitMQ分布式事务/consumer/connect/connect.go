package connect

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetRabbitMQ() (conn *amqp.Connection, err error) {
	var (
		IP       = "127.0.0.1"
		Port     = "5672"
		Account  = "admin"
		Password = "admin"
	)
	dataSource := fmt.Sprintf("amqp://%s:%s@%s:%s/", Account, Password, IP, Port)
	//获取连接
	conn, err = amqp.Dial(dataSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conn, nil
}

func GetMysql() (conn *gorm.DB, err error) {
	var (
		user     = "root"
		password = "123456"
		ip       = "127.0.0.1"
		port     = "3306"
		db       = "order_server"
		charset  = "utf8mb4"
	)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, password, ip, port, db, charset)
	conn, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return conn, nil
}
