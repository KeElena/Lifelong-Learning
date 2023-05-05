package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取web服务对象
	server := gin.Default()
	//加载html页面
	server.LoadHTMLGlob("./templates/*")
	//初始化session
	//设置私钥
	store := cookie.NewStore([]byte("my private key"))
	//装载session
	server.Use(sessions.Sessions("mysession", store))
	//设置session数据
	server.GET("/set", func(context *gin.Context) {
		//获取session对象
		session := sessions.Default(context)
		//设置键值对
		session.Set("msg", "hello world")
		//设置过期时间
		session.Options(sessions.Options{MaxAge: 3600})
		//保存session数据
		err := session.Save()
		if err != nil {
			context.HTML(200, "index.html", gin.H{"msg": "session save error"})
		}
	})

	server.GET("/get", func(context *gin.Context) {
		session := sessions.Default(context)
		val := session.Get("msg")
		context.HTML(200, "index.html", gin.H{"msg": val})
	})
	//运行web服务
	server.Run("0.0.0.0:8081")
}
