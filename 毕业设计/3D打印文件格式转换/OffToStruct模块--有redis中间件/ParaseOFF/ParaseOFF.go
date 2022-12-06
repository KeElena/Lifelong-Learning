package ParaseOFF

import (
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

func GetVertex(vertexNum int, reader *bufio.Reader) (map[string]string, error) {
	//使用map表按顺序存储坐标
	m := make(map[string]string, vertexNum)
	//遍历坐标部分
	for i := 0; i < vertexNum; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil, err
		}
		//以字符串的形式存储顺序
		m[fmt.Sprintf("%d", i)] = string(line)
	}
	return m, nil
}

func splitFunc(str string) (strArr []string) {
	//以空格进行拆分值返回字符串数组
	strArr = strings.Fields(str)
	return
}

func StoreFace(m map[string]string, faceNum int, reader *bufio.Reader, redisDB *redis.Client, fileName string) error {
	//使用管道发送redis请求
	pipe := redisDB.Pipeline()
	defer pipe.Close()
	for i := 0; i < faceNum; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		temps := strings.Fields(string(line))
		//根据坐标顺序在表中获取坐标并存储
		err = pipe.LPush(fileName, splitFunc(m[temps[1]])).Err()
		err = pipe.LPush(fileName, splitFunc(m[temps[2]])).Err()
		err = pipe.LPush(fileName, splitFunc(m[temps[3]])).Err()
		if err != nil {
			return err
		}
	}
	//发送redis请求
	pipe.Expire(fileName, time.Second)
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}
