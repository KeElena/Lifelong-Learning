package main

import (
	"encoding/json"
	"gin_jwt/JWT"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)
import "github.com/gin-contrib/static"

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	ExpireAt int64  `json:"expireAt"`
}

type Token struct {
	Uid      int64  `json:"uid"`
	Name     string `json:"name"`
	ExpireAt int64  `json:"expireAt"`
}

var key = []byte("123456")

func checkToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		payload, ok, err := JWT.Check(token, key)
		if err != nil {
			context.JSON(200, gin.H{
				"status": 200,
				"msg":    err.Error(),
			})
			context.Abort()
			return
		}
		//签名校验失败
		if !ok {
			context.JSON(200, gin.H{
				"status": 200,
				"msg":    "token解析失败，重新登录",
			})
			context.Abort()
			return
		}
		var auth Token
		err = json.Unmarshal(payload, &auth)
		if err != nil {
			context.JSON(200, gin.H{
				"status": 200,
				"msg":    "token解析失败，重新登录",
			})
			context.Abort()
			return
		}
		//token过期情况
		if auth.ExpireAt > time.Now().Unix() {
			context.Next()
			return
		} else {
			context.JSON(200, gin.H{
				"status": 200,
				"msg":    "token过期，重新登录",
			})
			context.Abort()
			return
		}
	}
}

func main() {
	server := gin.Default()
	server.Use(static.Serve("/", static.LocalFile("dist", true)))
	server.POST("/login", func(context *gin.Context) {
		var user User
		err := context.Bind(&user)
		if err != nil {
			log.Println(err)
			context.JSON(200, gin.H{"msg": "error"})
			return
		}
		if user.Account == "demo" && user.Password == "123456" {
			payload := Token{Uid: 1, Name: "demo", ExpireAt: time.Now().Add(time.Second * 5).Unix()}
			payloadByte, _ := json.Marshal(&payload)
			token := JWT.GetToken(payloadByte, key, JWT.SHA256)
			context.JSON(200, gin.H{
				"status": 200,
				"msg":    "登录成功",
				"token":  token,
			})
		} else {
			context.JSON(200, gin.H{
				"status": 200,
				"msg":    "匹配错误",
				"data":   nil,
			})
		}
	})

	server.GET("/get", checkToken(), func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": 200,
			"msg":    "请求处理成功",
		})
	})

	server.Run("0.0.0.0:8080")
}
