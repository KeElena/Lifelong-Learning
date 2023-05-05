package connect

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	grpcAddr = "127.0.0.1:9000"
)

var (
	mysqlUser    = "root"
	mysqlPwd     = "123456"
	mysqlIp      = "127.0.0.1"
	mysqlPort    = "3306"
	mysqlDB      = "eth"
	mysqlCharset = "utf8mb4"
)

var (
	mongoAddr = "mongodb://127.0.0.1:27017"
)

var (
	redisAddr = "127.0.0.1:6379"
)

func GRPC() (*grpc.ClientConn, error) {
	grpcConn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return grpcConn, nil
}

func MYSQL() (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", mysqlUser, mysqlPwd, mysqlIp, mysqlPort, mysqlDB, mysqlCharset)
	mysqlConn, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return mysqlConn, nil
}

func Mongo() (*mongo.Client, error) {
	option := options.Client().ApplyURI(mongoAddr)
	//可设置连接池
	option.SetMaxPoolSize(10)
	//获取连接
	client, err := mongo.Connect(context.Background(), option)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Redis() (*redis.Client, error) {
	redisConn := redis.NewClient(&redis.Options{Addr: redisAddr, Password: "", DB: 1})
	err := redisConn.Ping().Err()
	if err != nil {
		return nil, err
	}
	return redisConn, err
}
