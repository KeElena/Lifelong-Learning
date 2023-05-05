package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//获取web服务对象
	server := gin.Default()
	//加载html页面
	server.LoadHTMLGlob("./templates/*")
	server.GET("/set", func(context *gin.Context) {
		val, err := context.Cookie("msg")
		if err != nil {
			context.SetCookie("msg", "hello world", 3600, "/", "127.0.0.1", false, false)
			val = "hello world"
		}
		context.HTML(200, "index.html", gin.H{"msg": val})
	})
	//返回html页面
	server.GET("/get", func(context *gin.Context) {
		val, err := context.Cookie("msg")
		if err != nil {
			context.HTML(200, "index.html", gin.H{"msg": "error"})
		} else {
			context.HTML(200, "index.html", gin.H{"msg": val})
		}
	})

	//运行web服务
	server.Run("0.0.0.0:8081")
}
