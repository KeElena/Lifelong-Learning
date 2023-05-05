### 1. gRPC服务端的启动流程



### 2. gRPC服务类型有哪些

* 简单rpc：一个请求对象对应一个返回对象
* 服务端流式rpc：一个请求对象，服务端返回对公结果对象
* 客户端流式rpc：客户端发送发送多个请求对象，服务端返回一个响应结果
* 双向流式rpc：结合客户端流式rpc和服务端流式rpc，传入多个对象返回多个响应对象

[grpc的四种服务类型](https://www.cnblogs.com/resentment/p/6792029.html)

### 3. 一个connection可以同时处理多个steam，那keepalive是针对Steam设置，还是针对connection



### 4. http2 conn的keepalive处理流程



### 5. gRPC 通信报文格式

* 常见的使用`protobuf`的编码格式，可以是结构化的数据串行化
* 使用`protobuf`序列化后会得到二进制形式的数据，其占用的存储空间比json和xml小，适用于网络传输

### 6. 常见的拦截器有哪些，用开源库实现 还是自实现



### 7. 多路复用指的是什么



### 8. 传输报文中metadata通常存放哪些内容



### 9. 如何自定义resolver



### 10. 如何自定义balance



### 11. 如何实现gRPC全链路追踪



### 12. 客户端connection 连接状态有哪些



### 13. 客户端如何拿到服务端的服务函数List



### 14. gRPC 什么是backoff协议



### 15. gRPC如何为每个Steam进行限流，什么是Flow Control



### 16. 什么是HPack

