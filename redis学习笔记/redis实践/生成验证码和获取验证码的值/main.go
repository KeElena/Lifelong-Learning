func verifyCode() {
	var err error
	var val interface{}
	//随机生成验证码并暂存到redis
	rand.Seed(time.Now().UnixNano())
	err = redisDB.Set("code", rand.Int31n(999999), time.Second*12).Err()
	if err != nil {
		fmt.Println(err) 
		return
	}
	//redis里获取验证码
	val, err = redisDB.Get("code").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}