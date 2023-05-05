package main

import (
	"github.com/gin-gonic/gin"
)

//设置中间件
func myHander() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取id数据
		id := context.Param("id")
		//判断拦截请求
		if id == "1" {
			context.Abort()
		}
		//设值
		context.Set("msg", "access")
		//请求通过
		context.Next()
	}
}

func main() {
	//获取web服务对象
	server := gin.Default()
	//注册中间件
	server.Use(myHander())
	//返回json数据
	server.GET("/:id", func(context *gin.Context) {
		//获取中间件设置的数据
		msg := context.MustGet("msg").(string)
		//返回数据
		context.JSON(200, gin.H{"msg": msg})
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
