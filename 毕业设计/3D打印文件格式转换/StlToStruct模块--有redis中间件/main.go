package main

import (
	StlPkg "StlToStruct/ParseSTL"
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"net"
	"os"
	"time"
)

var redisDB *redis.Client

func ReadFile(fileObj *os.File) (err error) {
	//自定义缓冲
	reader := bufio.NewReaderSize(fileObj, 10*1024*1024)
	//使用管道处理
	pipe := redisDB.Pipeline()
	defer pipe.Close()
	for {
		//固定50个字节循环读取
		BufData := make([]byte, 50)
		n, _ := reader.Read(BufData)
		if n == 0 {
			break
		}
		//获取面单元
		unit := StlPkg.GetUnit(BufData)

		A := make([]string, 3)
		B := make([]string, 3)
		C := make([]string, 3)
		for i := 0; i < 3; i++ {
			A[i] = fmt.Sprintf("%f", unit.VertexA[i])
			B[i] = fmt.Sprintf("%f", unit.VertexB[i])
			C[i] = fmt.Sprintf("%f", unit.VertexC[i])
		}
		pipe.LPush(fileObj.Name(), A)
		pipe.LPush(fileObj.Name(), B)
		pipe.LPush(fileObj.Name(), C)
	}
	pipe.Expire(fileObj.Name(), time.Second)
	_, err = pipe.Exec()
	return err
}

//初始化redis对象
func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0})
	_, err = redisDB.Ping().Result()
	return
}

func transform(filePath string) (err error) {
	//打开文件
	fileObj, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败！")
		return
	}
	defer fileObj.Close()
	//跳过前84个字节
	_, err = fileObj.Seek(84, 0)
	if err != nil {
		fmt.Println("文件指针移动失败!")
		return
	}
	//读取文件内容
	err = ReadFile(fileObj)
	if err != nil {
		fmt.Println("redis数据存储失败")
		return
	}
	return
}

func handleFunc(conn net.Conn) {
	filePath := make([]byte, 256)
	var n int
	var err error
	//读取文件路径
	n, err = conn.Read(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//转换方法
	err = transform(string(filePath[:n]))
	//返回工作状态
	if err != nil {
		_, err = conn.Write([]byte("read file fail"))
		return
	}
	_, err = conn.Write([]byte("OK"))
	//状态返回失败时退出
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
	//构建tcp服务端
	listener, err = net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取tcp请求
	for {
		//获取tcp请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//处理请求
		handleFunc(conn)
		//关闭请求
		_ = conn.Close()
	}
}
