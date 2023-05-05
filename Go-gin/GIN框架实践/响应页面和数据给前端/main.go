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
	server.Static("/static", "./static")
	//构建一个GET请求
	server.GET("/hello", func(context *gin.Context) {
		//返回页面和相关数据
		context.HTML(200, "index.html", gin.H{
			"msg": "hello MyWeb",
		})
	})

	//运行web服务
	server.Run("0.0.0.0:8081")
}
