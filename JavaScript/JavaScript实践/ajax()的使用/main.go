package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("./templates/*")
	server.Static("/static", "./static")

	server.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})

	server.GET("/getValue", func(context *gin.Context) {
		v := context.Query("msg")
		if len(v) == 0 {
			context.JSON(200, gin.H{"msg": "null"})
		}
		context.JSON(200, gin.H{"msg": "hello " + v})
	})
	server.Run("0.0.0.0:8080")
}
