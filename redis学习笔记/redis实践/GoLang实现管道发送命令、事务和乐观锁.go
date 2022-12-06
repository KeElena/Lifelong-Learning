package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisDB *redis.Client

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0})
	_, err = redisDB.Ping().Result()
	return
}
//管道
func pipeFunc() {
	pipe := redisDB.Pipeline()
	defer pipe.Close()

	val1 := pipe.Set("test", "test", 0) 		//错误命令，用于理解区分是否有tx（事务）
	val2 := pipe.Incr("test")					//命令错误时事务锁不会执行的，但非事务会继续执行

	pipe.Exec()									//发送命令

	if val1.Err() != nil {
		fmt.Println(val1.Err())
	}
	fmt.Println(val1.Val())

	if val2.Err() != nil {
		fmt.Println(val2.Err())
	}
	fmt.Println(val2.String())
}
//事务、方法一
func workFuncOne() {
	//开启事务
	worker := redisDB.TxPipeline()
	defer worker.Close()
	//接收传回的cmd值
	val1 := worker.MSet([]string{"one", "1", "two", "2", "three", "3", "four", "4"})
	val2 := worker.Incr("one")
	val3 := worker.MGet([]string{"one", "two", "three", "four"}...)
	//执行事务
	_, _ = worker.Exec() //没有exec事务不会执行

	//判断是否有错，没有则输出值
	if val1.Err() != nil {
		fmt.Println(val1.Err())
	} else {
		fmt.Println(val1.Val())
	}
	if val2.Err() != nil {
		fmt.Println(val2.Err())
	} else {
		fmt.Println(val2.Val())
	}
	if val3.Err() != nil {
		fmt.Println(val3.Err())
	} else {
		fmt.Println(val3.Val())
	}
}
//事务方法二
func workFuncTwo() {
	var val1 *redis.StatusCmd
	var val2 *redis.IntCmd
	var val3 *redis.SliceCmd
	//使用闭包构造函数并构建命令
	fn := func(piper redis.Pipeliner) error {
		val1 = piper.MSet([]string{"one", "1", "two", "2", "three", "3", "four", "4"})
		val2 = piper.Incr("one")
		val3 = piper.MGet([]string{"one", "two", "three", "four"}...)
		return nil
	}
	//执行事务
	_, _ = redisDB.TxPipelined(fn)
	//判断是否有错，没有则输出值
	if val1.Err() != nil {
		fmt.Println(val1.Err())
	} else {
		fmt.Println(val1.Val())
	}
	if val2.Err() != nil {
		fmt.Println(val2.Err())
	} else {
		fmt.Println(val2.Val())
	}
	if val3.Err() != nil {
		fmt.Println(val3.Err())
	} else {
		fmt.Println(val3.Val())
	}
}
//乐观锁
func watchOP() {
	var val *redis.IntCmd
	var err error
	//使用闭包构建txf函数
	txf := func(tx *redis.Tx) (err error) {
		err = tx.Get("two").Err()
		if err != nil && err != redis.Nil {
			return err
		}

		//time.Sleep(time.Second * 6) //在命令行6秒内修改值用于测试乐观锁是否生效

		//使用闭包构建fn函数，val为txf函数所在环境的值（闭包的闭包）
		fn := func(piper redis.Pipeliner) error {
			val = piper.Incr("two")
			return nil
		}
		//构建事务
		_, err = tx.TxPipelined(fn)
		return
	}
	//使用乐观锁并执行事务
	err = redisDB.Watch(txf, "two")
	if err != nil {
		fmt.Println(err)
		return
	}
	if val.Err() != nil {
		fmt.Println(val.Err())
		return
	} else {
		fmt.Println(val.Val())
	}
}

func main() {
	var err error
	err = initRedis()
	if err != nil {
		fmt.Println(err)
		return
	}
	pipeFunc()
	//workFuncOne()
	//workFuncTwo()
	//watchOP()
}
