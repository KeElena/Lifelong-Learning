package main

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Msg string `json:"msg"`
}

func main() {
	//获取web服务对象
	server := gin.Default()
	//加载html页面
	server.LoadHTMLGlob("./templates/*")
	//返回html页面
	server.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	//返回字符串
	server.GET("str", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	//返回JSON数据
	//gin.H{}返回JSON数据
	server.GET("json1", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello world"})
	})
	//散列表返回JSON数据
	server.GET("json2", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg": "hello world",
		})
	})
	//结构体返回JSON数据
	server.GET("json3", func(context *gin.Context) {
		msg := &Message{Msg: "hello world"}
		context.JSON(200, msg)
	})
	//返回JSON跨域数据
	server.GET("jsonp", func(context *gin.Context) {
		context.JSONP(200, gin.H{"msg": "hello world"})
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
