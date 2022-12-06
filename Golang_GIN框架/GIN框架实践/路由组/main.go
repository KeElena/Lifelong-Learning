package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	//获取web服务对象
	server := gin.Default()
	//加载中间件---加载标签页icon
	server.Use(favicon.New("./刻晴-夜宵.png"))
	//加载资源
	server.LoadHTMLGlob("templates/*")
	//路由组的设计
	usrGroup := server.Group("/usr")
	{
		usrGroup.GET("/login", func(context *gin.Context) {
			context.String(200, "login")
		})
		usrGroup.GET("/register", func(context *gin.Context) {
			context.String(200, "register")
		})
	}
	//运行web服务
	server.Run("0.0.0.0:8081")
}
