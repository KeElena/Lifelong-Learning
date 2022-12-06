package main

import (
	"DxfToStruct/ParaseOFF"
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"net"
	"os"
)

var redisDB *redis.Client

//初始化redis对象
func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0})
	_, err = redisDB.Ping().Result()
	return
}

func transform(filePath string) error {
	//打开文件
	fileObj, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileObj.Close()
	//获取bufio对象
	reader := bufio.NewReader(fileObj)

	for {
		temp, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		if string(temp) == "ENTITIES" {
			break
		}
	}

	//根据坐标存储面
	err = ParaseOFF.GetFace(reader, redisDB, fileObj.Name())
	if err != nil {
		return err
	}
	return nil
}

func handleFunc(conn net.Conn) {
	var err error
	var n int
	filePath := make([]byte, 256)

	//读取文件路径
	n, err = conn.Read(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//进行转换操作
	err = transform(string(filePath[:n]))
	if err != nil {
		//转换操作失败返回失败消息
		_, _ = conn.Write([]byte("transform fail"))
		return
	}

	//转换成功则返回OK
	_, err = conn.Write([]byte("OK"))
	if err != nil {
		return
	}
}

func main() {
	var err error
	var listener net.Listener
	//初始化redis
	err = initRedis()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer redisDB.Close()
	//初始化tcp服务端
	listener, err = net.Listen("tcp", "0.0.0.0:8084")
	if err != nil {
		fmt.Println(err)
		return
	}
	//for循环处理请求
	for {
		//获取请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//处理请求
		handleFunc(conn)
		conn.Close()
	}

}
