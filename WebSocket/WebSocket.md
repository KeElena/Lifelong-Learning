# WebSocket

**一、WebSocket**

* `WebSocket`是 HTML5 开始提供的一种在单个 TCP 连接上进行全双工通讯的协议
* WebSocket 使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
* 浏览器和服务器只需要连接一次，两者之间就直接可以创建持久性的连接，并进行双向数据传输

**二、使用场景**

* 实时数据更新
* 大文件传输

**三、传统轮询推送**

* 轮询是在特定的的时间间隔由客户端向服务器发出HTTP请求，然后由服务器返回最新的数据给客户端
* 缺点：
  * 客户端需要不断的向服务器发出请求，容易引起`ddos`
  * HTTP请求可能包含较长的头部，其中有效的数据可能只是很小的一部分，浪费了带宽等资源

# JavaScript---客户端

**一、创建WebSocket对象**

* 使用构造方法`new WebSocket()`创建对象，需要输入`WebSockerUrl`进行初始化
* 创建socket对象时，会向url进行Ping操作，如果Ping失败则websocket对象的状态直接为close
* `WebSockerUrl`是使用`WebSocket`协议的连接地址，格式：`ws://IP:端口/路径`
* <font color=red>请求方式默认为`GET`</color>
* WebSocket需要发送一次请求才能与后端建立TCP连接，该过程可以放置一些数据

```javascript
let ws = new WebSocket("ws://localhost:9998/echo");
```

**二、WebSocket回调方法**

| 事件    | 回调方法名          | 描述                       |
| ------- | ------------------- | -------------------------- |
| open    | SocketOBJ.onopen    | 连接建立时触发             |
| message | SocketOBJ.onmessage | 客户端接收服务端数据时触发 |
| error   | SocketOBJ.onerror   | 通信发生错误时触发         |
| close   | SocketOBJ.onclose   | 连接关闭时触发             |

```javascript
this.webSocket = new WebSocket(url)
this.webSocket.onopen =function(){...}
this.webSocket.onclose =function(){...}
this.webSocket.onmessage =function(data){...}
this.webSocket.onerror =function(error){...}
```

**三、WebSocket方法**

* 页面销毁时需要关闭连接

| 方法              | 描述     |
| ----------------- | -------- |
| SocketOBJ.send()  | 发送数据 |
| SocketOBJ.close() | 关闭连接 |

# Vue3原生Javascript的客户端简单实现

```vue
<script setup>
import {reactive} from "vue";
import {onBeforeUnmount} from "vue";
//初始化变量
let info=reactive({msg:null,status:""})
let url="ws:192.168.2.41:8081/websocket"
//获取webSocket对象
let webSocket=new WebSocket(url)
//ready时执行
webSocket.onopen=function () {
  info.status="ready"
}
//有数据时执行
webSocket.onmessage=function (event) {
  console.log(event.data)
  info.msg=event.data
}
//服务端发生错误时执行
webSocket.onerror=function (error) {
  info.status="error:"+error
}
//关闭后执行
webSocket.onclose=function (){
  info.status="closed"
}
//发送消息以及建立WebSocket连接
function send(){
  webSocket.send("hello")
}
//自发停止函数
function stop(){
  webSocket.close()
  info.status="closed"
}

onBeforeUnmount(()=>{
  webSocket.close()
})
</script>
```

# Go-WebSocket

**一、导入第三方包**

* [github.com/gorilla/websocket文档](https://pkg.go.dev/github.com/gorilla/websocket)

```bash
go get github.com/gorilla/websocket
```

**二、升级连接对象**

* 将http端连接升级为TCP长连接前需要配置升级连接对象

| 参数              | 作用                                                         |
| ----------------- | ------------------------------------------------------------ |
| ReadBufferSize    | 读缓冲                                                       |
| WriteBufferSize   | 写缓冲                                                       |
| HandshakeTimeout  | 连接握手超时                                                 |
| CheckOrigin       | 检查header，过滤请求（true:通过，false:过滤）                |
| EnableCompression | 开启压缩（[RFC 7692](https://rfc-editor.org/rfc/rfc7692.html)），可以主动关闭 |

```go
var upgrader=websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    HandshakeTimeout: time.Second,
    CheckOrigin: func(r *http.Request)bool{
        return true
    },
}
```

**三、升级短链接为长连接**

* 使用`upgrader.Upgrade()`实现
* 需要输入请求的写对象，读对象，可以输入响应头信息(`map[string]interface{}`)

```go
//gin示例
conn,err:=upgrader.Upgrade(context.Writer,context.Request,nil)
if err!=nil{
    log.Println(err)
    return
}
defer conn.Close()
```

**四、读取数据**

* 使用`conn.ReadMessage()`读取客户端发来的数据
* 使用`conn.ReadJSON()`直接解析客户端发来的Json数据
* 返回消息类型（1为text，0为二进制）、数据和error

```go
dataType,data,err:=conn.ReadMessage()
```

**五、发送数据**

* 使用`conn.WriteMessage()`发送数据
* 响应输入数据类型（1为text，0为二进制）和字节数据

```go
err=conn.WriteMessage(1,[]byte(fmt.Sprintf("%d",i)))
if err!=nil{
    log.Println(err)
    return
}
```

**六、循环读取数据**

* 使用`conn.NextReader()`监听前端发来的数据

```go
func ReadLoop(conn *websocket.Conn){
	for{
        //获取监听数据
    	_,io,err:=conn.NextReader()
        if err!=nil{
            log.Println(err)
            return
        }
        //读取数据
        data,err:=ioutil.ReadAll(io)
        if err!=nil{
            log.Println(err)
            return
        }
        //处理数据
        log.Println(string(data))
    }
}
```

**七、并发问题**

* 不超过一个 goroutine 调用写入方法（NextWriter、SetWriteDeadline、WriteMessage、 WriteJSON， EnableWriteCompression， SetCompressionLevel）
* 不超过一个 goroutine 调用读取方法（NextReader， SetReadDeadline， ReadMessage， ReadJSON， SetPongHandler， SetPingHandler）

# Go-gin ---api服务端

## 实时响应的实现

```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	server := gin.Default()
	server.GET("/websocket", func(context *gin.Context) {
		//建立长连接
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		//读取客户端发来的初始数据
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//处理初始数据
		log.Println(string(data))
		//循环返回实时数据
		i := 0
		for {
			//控制返回频率
			time.Sleep(time.Second)
			i++
			//返回数据
			err = conn.WriteMessage(msgType, []byte(fmt.Sprintf("%d", i)))
			if err != nil {
				log.Println(err)
				return
			}
		}
	})

	server.Run("0.0.0.0:8081")
}
```

## 全双工交互的实现

```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	server := gin.Default()
	server.GET("/websocket", func(context *gin.Context) {
		//建立长连接
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		//读取客户端发来的初始数据
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//处理初始数据
		log.Println(string(data))
		//监听并读取客户端发来的数据
		go ReadLoop(conn)
		//循环返回实时数据
		i := 0
		for {
			//控制返回频率
			time.Sleep(time.Second)
			i++
			//返回数据
			err = conn.WriteMessage(msgType, []byte(fmt.Sprintf("%d", i)))
			if err != nil {
				log.Println(err)
				return
			}
		}
	})

	server.Run("0.0.0.0:8081")
}

func ReadLoop(conn *websocket.Conn){
	for{
        //获取监听数据
    	_,io,err:=conn.NextReader()
        if err!=nil{
            log.Println(err)
            return
        }
        //读取数据
        data,err:=ioutil.ReadAll(io)
        if err!=nil{
            log.Println(err)
            return
        }
        //处理数据
        log.Println(string(data))
    }
}
```

# Go-zero ---api服务端

**一、api文件**

```apl
service websocket-api {
	@handler WebsocketHandler
	get /websocket returns ()
}
```

**二、修改处理websocket的handler文件**

* 修改`l.websocket()`方法的输入参数
* 去除冗余无效部分代码

```go
func WebsocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewWebsocketLogic(r.Context(), svcCtx)
		err := l.Websocket(w, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, nil)
		}
	}
}
```

**三、修改处理websocket的logic文件**

```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (l *WebsocketLogic) Websocket(w http.ResponseWriter, r *http.Request) (err error) {
	// todo: add your logic here and delete this line
	//建立长连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	//读取客户端发来的初始数据
	msgType, data, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	//处理初始数据
	log.Println(string(data))
	//监听并读取客户端发来的数据
	go ReadLoop(conn)
	//循环返回实时数据
	i := 0
	for {
		//控制返回频率
		time.Sleep(time.Second)
		i++
		//返回数据
		err = conn.WriteMessage(msgType, []byte(fmt.Sprintf("%d", i)))
		if err != nil {
			log.Println(err)
			return
		}
	}

}

func ReadLoop(conn *websocket.Conn) {
	for {
		_, io, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}
		data, err := ioutil.ReadAll(io)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(data))
	}
}
```





[(36条消息) nginx中的超时设置，请求超时、响应等待超时等_nginx设置延迟响应_半隐退状态的博客-CSDN博客](https://blog.csdn.net/m0_67900727/article/details/123472127)