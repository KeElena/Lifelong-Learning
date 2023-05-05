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
	//构建一个GET请求
	server.GET("/hello", func(context *gin.Context) {
		//返回JSON内容
		context.JSON(200, gin.H{"msg": "hello world"})
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
