package main

import (
	"BridgeModule/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	//log
	logFile := "./error.log"
	file, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	log.SetOutput(file)
	//获取gin对象
	webServer := gin.Default()
	//装载session中间件
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("my private key"))
	if err != nil {
		fmt.Println("redis error")
		return
	}
	//装载session
	webServer.Use(sessions.Sessions("userinfo", store))
	//装载路由
	router.InitRouter(webServer)
	webServer.Run("0.0.0.0:9090")
}
