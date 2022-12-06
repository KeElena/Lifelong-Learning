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

func StringType() {
	var val interface{}
	var err error
	//设置有时间限制的值
	err = redisDB.Set("myval", "2", time.Millisecond*2).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置永久的键值对
	err = redisDB.Set("myval", "1", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取值
	val, err = redisDB.Get("myval").Int()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//追加字符串
	err = redisDB.Append("myval", "1").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取字符串长度
	val = redisDB.StrLen("myval").Val()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("len:", val)
	//自增+1
	err = redisDB.Incr("myval").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//自减-1
	err = redisDB.Decr("myval").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//增加任意数字
	err = redisDB.IncrBy("myval", 10).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//截取字符串
	val, err = redisDB.GetRange("myval", 0, -1).Bytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//替换指定位置字符串
	err = redisDB.SetRange("myval", 0, "1").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//不存在则设置
	err = redisDB.SetNX("myval", "1", time.Second).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//一次存多个值
	err = redisDB.MSet([]string{"one", "1", "two", "2"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//一次取多个值
	vals := redisDB.MGet([]string{"one", "two"}...).Val()
	for _, v := range vals {
		fmt.Println(v.(string))
	}
	//一次存多个不存在的值
	err = redisDB.MSetNX([]string{"one", "1", "two", "2"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//取值并换值
	val, err = redisDB.GetSet("myval", "10").Int()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func ListType() {
	var err error
	var val interface{}
	//头部插入多个值
	err = redisDB.LPush("arr", []string{"one", "two", "three"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//头部取值
	val, err = redisDB.LRange("arr", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//尾部加值
	err = redisDB.RPush("arr", []string{"hello", "world"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//从左边开始获取并移除1个元素
	val, err = redisDB.LPop("arr").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//右边开始移除多个元素
	val, err = redisDB.RPop("arr").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//根据索引获取list的一个元素
	val, err = redisDB.LIndex("arr", 1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取list的长度
	val, err = redisDB.LLen("arr").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//根据值移除列表元素
	err = redisDB.LRem("arr", 2, "three").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//修减列表
	err = redisDB.LTrim("arr", 1, 16).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//移除一个列表的最后一个元素，将元素插到新列表的头部
	val, err = redisDB.RPopLPush("arr", "new").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//修改列表某个值
	err = redisDB.LSet("arr", 3, "world").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//在某个元素前插入新值
	err = redisDB.LInsert("arr", "before", "world", "hello").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SetType() {
	var err error
	var val interface{}
	//给集合添加值
	err = redisDB.SAdd("myset", []string{"one", "two", "three", "four", "five", "six", "seven"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//检测集合是否存在某元素
	val, err = redisDB.SIsMember("myset", "two").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取集合的长度
	val = redisDB.SCard("myset").Val()
	if err != nil {
		fmt.Println(err)
		return
	}
	//移除集合中某个元素
	val, err = redisDB.SRem("myset", "three").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//随机获取1个元素或多个元素
	val = redisDB.SRandMember("myset").Val()
	_, _ = redisDB.SRandMemberN("myset", 2).Result()
	fmt.Println(val)
	//获取集合所有元素
	val, err = redisDB.SMembers("myset").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//随机获取并删除集合多个值
	val, err = redisDB.SPop("myset").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//redisDB.SPopN("myset",2).Result()
	//移动集合元素岛另一个集合
	val, err = redisDB.SMove("myset", "other", "six").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//求交集
	val, err = redisDB.SInter("myset", "other").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//求差集
	val, err = redisDB.SDiff("myset", "other").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//求并集
	val, err = redisDB.SUnion("myset", "other").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}
func HashType() {
	var err error
	var val interface{}
	//添加1个键值对
	val, err = redisDB.HSet("myhash", "one", 1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//添加多个键值对
	m := make(map[string]interface{})
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	val, err = redisDB.HMSet("myhash", m).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取key里某个键值对
	val, err = redisDB.HGet("myhash", "one").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取key里所有键值对
	val, err = redisDB.HGetAll("myhash").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for key, v := range val.(map[string]string) {
		fmt.Println(key, "->", v)
	}
	//删除key里某个键值对
	err = redisDB.HDel(",myhash", "five").Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取hash表长度
	val, err = redisDB.HLen("myhash").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//判断hash表是否存在某个键值对
	val, err = redisDB.HExists("myhash", "two").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取hash表中所有key
	val, err = redisDB.HKeys("myhash").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//获取hash表中所有value
	val, err = redisDB.HVals("myhash").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//如果key中不存在这键值对则设置键值对
	val, err = redisDB.HSetNX("myhash", "hello", "world").Result()
	if err != nil {
		fmt.Println("key不存在", err)
	} else {
		fmt.Println(val)
	}
	//hash表中值的增加，有用于浮点数的方法，返回增加后的值
	val, err = redisDB.HIncrBy("myhash", "one", 1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func ZsetType() {
	var err error
	var val interface{}
	//设置有序集合元素
	temp := make([]redis.Z, 0, 5)
	temp = append(temp, redis.Z{Member: "one", Score: 1})
	temp = append(temp, redis.Z{Member: "two", Score: 2})
	temp = append(temp, redis.Z{Member: "three", Score: 3})
	temp = append(temp, redis.Z{Member: "four", Score: 4})
	temp = append(temp, redis.Z{Member: "five", Score: 5})
	err = redisDB.ZAdd("myset", temp...).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//按score从小到大获取范围内的元素
	val, err = redisDB.ZRange("myset", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//按score从大到小获取范围内的元素
	val, err = redisDB.ZRevRange("myset", 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//根据score的范围取值
	val, err = redisDB.ZRangeByScore("myset", redis.ZRangeBy{Min: "2", Max: "4"}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//移除元素
	err = redisDB.ZRem("myset", []string{"one", "two"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取有序集合岛元素个数
	val, err = redisDB.ZCard("myset").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取score区间里元素个数
	val, err = redisDB.ZCount("myset", "0", "5").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//增加有序集合score值
	val, err = redisDB.ZIncrBy("myset", 2, "five").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func GeospatialType() {
	var err error
	var val interface{}
	//添加地理位置
	temp := make([]*redis.GeoLocation, 0, 4)
	temp = append(temp, &redis.GeoLocation{Longitude: 114.08, Latitude: 22.54, Name: "shenzhen"})
	temp = append(temp, &redis.GeoLocation{Longitude: 113.12, Latitude: 23.02, Name: "foshan"})
	temp = append(temp, &redis.GeoLocation{Longitude: 113.28, Latitude: 23.12, Name: "guangzhou"})
	temp = append(temp, &redis.GeoLocation{Longitude: 121.47, Latitude: 31.23, Name: "shanghai"})
	err = redisDB.GeoAdd("city", temp...).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//查询存储的经度和纬度
	val, err = redisDB.GeoPos("city", []string{"foshan", "shenzhen"}...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]*redis.GeoPos) {
		fmt.Println(v)
	}
	//获取两个坐标间岛距离
	val, err = redisDB.GeoDist("city", "foshan", "shanghai", "km").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//获取某坐标一定半径内地理集合中的元素
	val, err = redisDB.GeoRadius("city", 112.97, 22.92, &redis.GeoRadiusQuery{Radius: 1000, Unit: "km"}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]redis.GeoLocation) {
		fmt.Println(v)
	}
	//获取地理集合里某元素一定半径内的所有元素
	val, err = redisDB.GeoRadiusByMember("city", "foshan", &redis.GeoRadiusQuery{Radius: 100, Unit: "km"}).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]redis.GeoLocation) {
		fmt.Println(v)
	}
	//获取坐标的hash编码
	val, err = redisDB.GeoHash("city", []string{"foshan", "shenzhen"}...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range val.([]string) {
		fmt.Println(v)
	}
	//删除地理元素
	err = redisDB.ZRem("city", []string{"foshan", "shenzhen"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func HyperloglogType() {
	var err error
	var val interface{}
	//构造PF集合并添加多个元素
	err = redisDB.PFAdd("myset", []string{"one", "two", "three", "four"}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//统计基数
	val, err = redisDB.PFCount("myset").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//合并PF集合
	val, err = redisDB.PFMerge("myset", "other").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func BitmapType() {
	var err error
	var val interface{}
	//设置值
	err = redisDB.SetBit("sign", 2, 1).Err()
	err = redisDB.SetBit("sign", 1, 1).Err()
	err = redisDB.SetBit("sign", 10, 1).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取值
	val, err = redisDB.GetBit("sign", 2).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	//统计1的个数
	val, err = redisDB.BitCount("sign", &redis.BitCount{}).Result() //BitCount可以为空
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func OtherOperation() {
	var val interface{}
	var err error
	//判断key是否存在
	val, err = redisDB.Exists("myval").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//转移key
	val, err = redisDB.Move("myval", 1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//设置过期时间
	val, err = redisDB.Expire("arr", time.Second*10).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	//查看数据距离过期剩余的时间
	val, err = redisDB.TTL("arr").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func main() {
	var err error
	err = initRedis()
	defer redisDB.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	StringType()
	ListType()
	//SetType()
	//HashType()
	//ZsetType()
	//GeospatialType()
	//HyperloglogType()
	//BitmapType()
	OtherOperation()
}
