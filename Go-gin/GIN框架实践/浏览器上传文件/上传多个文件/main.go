package main

import (
	"fmt"
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
	//返回上传html页面
	server.GET("/file", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	//获取上传的文件
	server.POST("/file", func(context *gin.Context) {
		//获取文件表单
		form, _ := context.MultipartForm()
		//根据key获取文件集合
		files := form.File["image"]
		//遍历每个gin的文件指针并保存数据
		for _, file := range files {
			_ = context.SaveUploadedFile(file, "./save/"+file.Filename)
		}
		context.String(200, fmt.Sprintf("文件上传成功！"))
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
