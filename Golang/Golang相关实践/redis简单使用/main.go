package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisDB *redis.Client

func initRedis() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisDB.Ping().Result()

	return
}

func setValue(key string, val int) error {
	err := redisDB.Set(key, val, time.Second*5).Err()
	return err
}

func getValue(key string) (val int, err error) {
	val, err = redisDB.Get(key).Int()
	return val, err
}

func sortedSet() (num int64, err error) {
	zsetKey := "Language_Rank"
	languages := []redis.Z{
		{Score: 10, Member: "Java"},
		{Score: 20, Member: "Python"},
		{Score: 30, Member: "PhP"},
		{Score: 40, Member: "C"},
		{Score: 50, Member: "Go"},
	}
	num, err = redisDB.ZAdd(zsetKey, languages...).Result()
	return
}

func addPoint(key string, addNum float64, elem string) (score float64, err error) {
	score, err = redisDB.ZIncrBy(key, addNum, elem).Result()
	return
}

func main() {

	var err error
	var val int
	//var num int64
	var score float64
	err = initRedis()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("OK")
	defer redisDB.Close()

	err = setValue("age", 5)
	if err != nil {
		fmt.Println(err)
		return
	}

	val, err = getValue("age")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v:%T\n", val, val)

	//num, err = sortedSet()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("num:", num)

	score, _ = addPoint("Language_Rank", 5, "Java")
	fmt.Println(score)
}
