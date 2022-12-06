package ParaseOFF

import (
	"bufio"
	"github.com/go-redis/redis"
	"time"
)

func GetFace(reader *bufio.Reader, redisDB *redis.Client, fileName string) error {
	//使用管道发送redis请求
	pipe := redisDB.Pipeline()
	defer pipe.Close()
	for {
		for {
			temp, _, err := reader.ReadLine()
			if err != nil {
				if err.Error() != "EOF" {
					return err
				} else {
					goto result
				}
			}

			if string(temp) == "3DFACE" {
				_, _, err = reader.ReadLine()
				_, _, err = reader.ReadLine()
				if err != nil {
					return err
				}
				break
			}
		}
		for i := 0; i < 18; i++ {
			if i%2 == 0 {
				_, _, _ = reader.ReadLine()
				continue
			}
			temp, _, _ := reader.ReadLine()
			pipe.LPush(fileName, string(temp))
		}
	}
result:
	//发送redis请求
	pipe.Expire(fileName, time.Second)
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}
