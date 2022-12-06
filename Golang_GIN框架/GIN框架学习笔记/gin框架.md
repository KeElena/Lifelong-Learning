# 下载Gin框架

**一、安装方式**

```shell
go get -u github.com/gin-gonic/gin
```

# Web服务的简单构建

**一、获取web服务对象**

* 可以使用`gin.Default()`和`gin.New()`获取web服务对象
* `gin.Default()`会打印请求响应的速度，常用于开发阶段
* `gin.New()`不会打印请求响应的速度，常用于上线阶段

```go
server:=gin.Default()
server:=gin.New()
```

**二、构建响应请求的方法**

* `server.GET()`用于处理`Get`请求
* `server.POST()`用于响应`Post`请求
* `server.GET()`使用格式：`server.GET("响应路径",响应方法)`
* `server.POST()`使用格式：`server.POST("响应路径",响应方法)`

```go
server.GET("/hello",func(context *gin.Context){
    context.JSON(200,gin.H{"msg":"hello world"})
})
```

**三、启动Web服务**

* 使用`server.Run()`启动服务
* 格式：`server.Run("IP:端口")`

```go
server.Run("0.0.0.0:8081")
```

**四、给标签页添加`icon`**

* 下载第三方包：`go get github.com/thinkerou/favicon`
* 通过`server.Use()`方法使用中间件加载`icon`
* 使用第三方包的`New()`方法加载`icon`
* 第三方包使用方法`favicon.New("icon路径")`

```go
server.Use(favicon.New("./刻晴-夜宵.png"))
```

# 前后端请求处理--RESTful API

**一、RESTful API**

* 允许网页在同一路径下处理增删改查的请求
* 传统的网页需要在不同的路径用`POST`处理删改查

| RESTful请求类型 | 常见用处 |
| --------------- | -------- |
| GET             | 查询数据 |
| POST            | 增加数据 |
| PUT             | 更新数据 |
| DELETE          | 删除数据 |

**二、使用`RESTful API`**

* 处理`GET`请求

```go
server.GET("/hello", func(context *gin.Context) {
	//返回JSON内容
	context.JSON(200, gin.H{"msg": "SELECT"})
})
```

* 处理`POST`请求

```go
server.POST("/hello", func(context *gin.Context) {
	context.JSON(200, gin.H{"msg": "ADD"})
})
```

* 处理`PUT`请求

```go
server.PUT("/hello", func(context *gin.Context) {

	context.JSON(200, gin.H{"msg": "UPDATA"})

})
```

* 处理`DELETE`请求

```go
server.DELETE("/hello", func(context *gin.Context) {

	context.JSON(200, gin.H{"msg": "DELETE"})

})
```

# 响应页面与数据给前端

**一、加载资源**

* 加载一个`HTML`页面：使用`server.LoadHTMLFiles()`加载
* 加载全部`HTML`页面：使用`server.LoadHTMLGlob()`加载
* 加载目录里全部的静态资源：使用`server.Static()`加载
* 加载静态资源的方法使用格式：`server.Static("网页路径","主机存放资源的根目录")`

```go
server.LoadHTMLFiles("./templates/index.html")
server.LoadHTMLGlob("./templates/*")
server.Static("/static","./static")
```

**二、状态码**

* 用于表示请求的处理情况
* 可以直接写数字

```go
StatusOK=200
StaticNotFound=404
```

**三、响应页面与数据**

* 使用`context.HTML()`进行响应页面

```go
server.Get("/hello",func(context *gin.Context){
    context.HTML(200,"index.hrml",gin.H{
        "msg":"hello MyWeb",
    })
})
```

# 前端传递参数到后端

## 传统的Get请求取参

* 传统`Get`请求参数在`url`中以`?`开头，以`&`拼接
* 使用`context.Query()`通过`key`取`value`

```go
server.GET("/hello",func(context *gin.Context){
    id:=context.Query("id")
    name:=context.Query("name")
    context.HTML(200,"index.html",gin.H{
        "id":id,
        "name":name,
    })
})
```

## Get请求RESTful形式取参

* `RESTful`形式的请求参数以网页相对路径为主，用`/:`定义的`key`，数据传来时对应位置为`value`
* 使用`context.Param()`获取值

```go
//定义含RESTful形式的url，/:后面跟着key
server.Get("/hello/:id/:name",func(context *gin.Context){
    id:=context.Param("id")
    name:=context.Param("name")
    context.HTML(200,"index.html",gin.H{
        "id":id,
        "name":name,
    })
})
```

## Post请求取参

* 使用`context.GetRawData()`获取`json`数据
* 使用`json.Unmarshal()`实现反序列化（数据转为键值对）

```go
server.POST("/usr", func(context *gin.Context) {
	data, _ := context.GetRawData()
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	context.JSON(200, m)
})
```

## 获取表单数据

* 使用`context.PostForm()`根据`key`获取`value`

```go
server.POST("hello",func(context *gin.Context){
    id:=context.PostForm("id")
    passwd:=context.PostForm("passwd")
    context.JSON()
})
```

# 重定向

* 使用`conrext.Redirect()`实现网站重定向

```go
server.GET("/redirect",func(context *gin.Context){
    context.Redirect(301,"https://www.baidu.com")
})
```

# 自定义默认页面

## 自定义404页面

* 使用`server.NoRoute()`响应自定义的404页面

```go
server.NoRoute(func(context *gin.Context) {
	context.HTML(404, "404.html", nil)
})
```

# 路由组

**一、路由组**

* 实现路径的拼接
* 使用`server.Group()`创建并定义路由组的根路径

* 格式：`路由组对象:=server.Group("网页路径")`

* 可以使用`{ }`来框住路由组下的所有路径，提高可读性

```go
usrGroup := server.Group("/usr")
{
	usrGroup.GET("/login", func(context *gin.Context) {
		context.String(200, "login")
	})
	usrGroup.GET("/register", func(context *gin.Context) {
		context.String(200, "register")
	})
}
```

# 文件上传

## 文件上传非必要的设值

**一、限制文件大小**

* 使用`server.MaxMultipartMemory`设置 

```go
server.MaxMultipartMemory = 8 << 20		//限制8Mb  1MB
```

## 上传一个文件

* 使用`context.FormFile()`获取web文件指针
* 格式：`webfile,err:=context.FormFile("HTML key")`
* 使用`context.SaveUploadedFile()`保持文件
* 格式：`context.SaveUploadedFile(webfile,"目录路径/"+File.Filename)`
* webfile指针的方法：
  * `Filename()`：获取文件名
  * `Size()`：获取文件大小
  * `Open()`：获取文件指针`*File`

```go
server.POST("/file", func(context *gin.Context) {
    //获取web文件指针
	file, _ := context.FormFile("image")
	//打印日志
    log.Println("filename:", file.Filename)
	log.Println("size:", file.Size)
	//保存文件
	_ = context.SaveUploadedFile(file,"./save/+file.Filename")
    //返回处理信息
	context.String(200, fmt.Sprintf("%s上传成功！", file.Filename))
})
```

## 上传多个文件

* 使用`context.MultipartForm()`获取文件表单
* 使用`form.File["key"]`获取所有文件的web指针
* 循环遍历web指针并保存文件

```go
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
```

# 后端返回文件

* 使用`context.File()`返回文件

```go
server.GET("/zip", func(context *gin.Context) {
	context.File("./111.zip")
})
```













