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
	//限制上传的文件大小 8Mb
	server.MaxMultipartMemory = 8 << 20
	//返回上传html页面
	server.GET("/file", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	//返回文件
	server.GET("/zip", func(context *gin.Context) {
		context.File("./111.zip")
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
