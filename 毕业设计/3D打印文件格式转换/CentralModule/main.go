package main

import (
	"fmt"
	"format/GetFile"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/thinkerou/favicon"
	"net"
	"os"
	"os/exec"
	"strings"
)

var redisDB *redis.Client

func initWeb() *gin.Engine {
	server := gin.Default()
	server.Use(favicon.New("./kq.png"))
	server.LoadHTMLGlob("templates/*")
	return server
}

func initServer() (err error) {
	//初始化redis对象
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "myredis:6379",
		Password: "",
		DB:       0})
	_, err = redisDB.Ping().Result()
	if err != nil {
		return
	}
	//初始化StlToStruct模块
	err = exec.Command("./StlToStruct", "").Start()
	if err != nil {
		return
	}
	//初始化StlToStruct模块
	err = exec.Command("./OffToStruct", "").Start()
	if err != nil {
		return
	}
	err = exec.Command("./DxfToStruct", "").Start()
	if err != nil {
		return
	}
	err = exec.Command("./ObjToStruct", "").Start()
	if err != nil {
		return
	}
	return nil
}

func transform(filePath string, address string) string {
	var n int
	msg := make([]byte, 256)
	//发送tcp请求，UDP更优
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	//发送文件路径
	_, err = conn.Write([]byte(filePath))
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	//获取文件处理消息
	n, err = conn.Read(msg)
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	//返回处理消息
	return string(msg[:n])
}

func main() {
	//初始化其他模块服务
	err := initServer()
	if err != nil {
		panic(err)
	}
	//初始化web服务
	webServer := initWeb()

	webServer.GET("/file", func(context *gin.Context) {
		context.HTML(200, "web.html", nil)
	})

	webServer.POST("/hander", func(context *gin.Context) {
		var tip string
		//接收文件
		uploadFile, _ := context.FormFile("image")
		//限制文件大小
		if uploadFile.Size > 5*1024*1024 {
			context.HTML(200, "error.html", gin.H{"error": "文件过大"})
			return
		}
		if len([]byte(uploadFile.Filename)) >= 256 {
			context.HTML(200, "error.html", gin.H{"error": "文件名称过长"})
			return
		}
		//文件暂存
		_ = context.SaveUploadedFile(uploadFile, "./temp/"+uploadFile.Filename)
		//文件分类处理
		if strings.HasSuffix(uploadFile.Filename, ".stl") {
			//根据后缀发送文件路径到不同模块，获取处理消息
			tip = transform("./temp/"+uploadFile.Filename, "0.0.0.0:8082")
		} else if strings.HasSuffix(uploadFile.Filename, ".dxf") {
			tip = transform("./temp/"+uploadFile.Filename, "0.0.0.0:8084")
		} else if strings.HasSuffix(uploadFile.Filename, ".off") {
			tip = transform("./temp/"+uploadFile.Filename, "0.0.0.0:8083")
		} else if strings.HasSuffix(uploadFile.Filename, ".obj") {
			tip = transform("./temp/"+uploadFile.Filename, "0.0.0.0:8085")
		}

		//处理转换错误
		if tip != "OK" {
			context.HTML(200, "error.html", gin.H{"error": tip})
			return
		}
		//返回文件
		//设置响应头参数
		context.Header("Content-Transfer-Encoding", "binary")
		context.Header("Cache-Control", "no-cache")
		context.Header("Content-Type", "application/octet-stream")
		//判断返回文件类型
		err = GetFile.GetFile(redisDB, uploadFile.Filename)
		if err != nil {
			context.HTML(200, "error.html", gin.H{"error": "File parsing failed!"})
		}

		if !strings.HasSuffix(uploadFile.Filename, ".dxf") {
			context.Header("Content-Disposition", "attachment;filename="+uploadFile.Filename[:len(uploadFile.Filename)-4]+".dxf")
			context.File("./save/" + uploadFile.Filename[:len(uploadFile.Filename)-4] + ".dxf")
			_ = os.Remove("./save/" + uploadFile.Filename[:len(uploadFile.Filename)-4] + ".dxf")
		} else {
			context.Header("Content-Disposition", "attachment;filename="+uploadFile.Filename[:len(uploadFile.Filename)-4]+".stl")
			context.File("./save/" + uploadFile.Filename[:len(uploadFile.Filename)-4] + ".stl")
			_ = os.Remove("./save/" + uploadFile.Filename[:len(uploadFile.Filename)-4] + ".stl")
		}
		//删除文件
		_ = os.Remove("./temp/" + uploadFile.Filename)
	})
	_ = webServer.Run("0.0.0.0:8081")
}
