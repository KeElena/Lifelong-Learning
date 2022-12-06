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

	server.GET("/redirect", func(context *gin.Context) {
		context.Redirect(301, "https://www.baidu.com")
	})

	//运行web服务
	server.Run("0.0.0.0:8081")
}
