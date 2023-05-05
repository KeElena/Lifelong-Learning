package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//获取web服务对象
	server := gin.Default()
	//加载资源
	server.LoadHTMLGlob("templates/*")
	//返回html页面
	server.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	//限制文件相关参数
	server.POST("file", func(context *gin.Context) {
		//接收文件
		uploadFile, _ := context.FormFile("image")
		//限制文件大小
		if uploadFile.Size > 6*1024*1024 {
			context.HTML(200, "msg.html", gin.H{"msg": "文件过大"})
			return
		}
		//限制文件名长度
		if len([]byte(uploadFile.Filename)) >= 256 {
			context.HTML(200, "msg.html", gin.H{"msg": "文件名称过长"})
			return
		}
		//文件暂存
		//_ = context.SaveUploadedFile(uploadFile, "./temp/"+uploadFile.Filename)

		context.HTML(200, "msg.html", gin.H{"msg": "上传成功"})
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
