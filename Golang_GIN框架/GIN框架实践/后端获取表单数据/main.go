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
	//传统Get请求取参
	server.GET("/hello", func(context *gin.Context) {

		context.HTML(200, "index.html", nil)
	})

	server.POST("/hello", func(context *gin.Context) {
		id := context.PostForm("id")
		passwd := context.PostForm("passwd")
		context.JSON(200, gin.H{
			"id":     id,
			"passwd": passwd,
		})
	})

	//运行web服务
	server.Run("0.0.0.0:8081")
}
