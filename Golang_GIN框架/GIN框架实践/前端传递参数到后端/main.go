package main

import (
	"encoding/json"
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
		id := context.Query("id")
		name := context.Query("name")
		//返回页面和相关数据
		context.HTML(200, "index.html", gin.H{
			"id":   id,
			"name": name,
		})
	})
	//Get请求RESTful形式取参
	server.GET("/hello/:id/:name", func(context *gin.Context) {
		id := context.Param("id")
		name := context.Param("name")
		context.HTML(200, "index.html", gin.H{
			"id":   id,
			"name": name,
		})
	})
	//Post请求取参
	server.POST("/usr", func(context *gin.Context) {
		data, _ := context.GetRawData()
		var m map[string]interface{}
		json.Unmarshal(data, &m)
		context.JSON(200, m)
	})

	//运行web服务
	server.Run("0.0.0.0:8081")
}
