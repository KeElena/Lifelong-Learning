# MongoDB简介

**一、MongoDB**

* MongoDB是一个开源、高性能、无模式的文档型数据库
* MongoDB中记录的是一个文档，由字段和键值对组成的数据结构
* MongoDB文档类似JSON对象，字段的数据类型是字符型
* MongoDB不支持表连接

**二、MongoDB相关概念**

| 概念        | 说明                           |
| ----------- | ------------------------------ |
| database    | 数据库                         |
| collection  | 集合                           |
| document    | 文档                           |
| field       | 域                             |
| index       | 索引                           |
| 嵌入文档    | 通过嵌入式文档代替多表连接     |
| primary key | MongoDB字段将_id字段设置为主键 |

**三、数据类型**

* MongoDB使用`BSON`存储数据

| 数据类型   | 描述                               | 示例                    |
| ---------- | ---------------------------------- | ----------------------- |
| 字符串     | UTF-8字符串数据                    | {"key":"hello"}         |
| 对象id     | 对象id是文档的唯一id，用12字节表示 | {"key":"ObjectId()"}    |
| 布尔值     | true和false                        | {"key":"true"}          |
| 数组       | 数组                               | {"key":"["a","b","c"]"} |
| 64位浮点数 | 浮点数                             | {"x":3.1415,"y":3}      |
| null       | 空                                 | {"key":"null"}          |
| undefined  | 未定义的类型                       | {"key":"undefined"}     |
| 正则表达式 | 正则语法                           | {"key":正则表达式}      |
| 代码       | JavaScript代码                     | {"key":function{/**/}}  |

**三、MongoDB的特点**

* 高性能
  * 提供高性能的数据持久性
  * 使用嵌入式数据模型减少了数据库系统的I/O次数
* 高可用
  * 提供自动故障转移和数据冗余
* 高扩展

**四、应用场景**

* 社交场景：使用MongoDB存储用户信息以及用户发表的动态，实现附近人，地点等功能
* 游戏场景：存储用户信息，用户的装备、积分，以内嵌文档的形式存储，实现高效存储和访问
* 物流场景：使用MongoDB存储订单信息，订单状态在运送过程中不断更新，以内嵌数组形式存储
* 物联网场景：使用MongoDB存储所有接入智能设备信息，以及设备的日志信息
* 视频直播：所有MongoDB存储用户信息、点赞互动信息等

**五、应用场景普遍特点**

* 数据量大
* 读写操作频繁
* 价值较低，对事务性要求不高

# docker中安装MongoDB

**一、拉取镜像**

```shell
docker pull mongo
```

**二、启动容器**

* mongodb容器默认没有身份认证

```shell
docker run -d 									//后台运行
			-p 27017:27017 						//端口映射
			-v mongo_configdb:/data/configdb 	//挂载配置文件目录
			-v mongo_db:/data/db 				//挂载数据目录
			--name mymongo 						//容器名
			mongo								//使用的镜像
			--auth								//启动身份认证，可选
```

**三、创建管理员账号**

* 使用了--auth时

```shell
//进入容器
docker exec -it mymongo mongo admin
```

```bash
//创建管理员账号
db.createUser({ user: 'root', pwd: '123456', roles: [ { role: "userAdminAnyDatabase", db: "admin" },"readWriteAnyDatabase" ] });
```

**四、其他操作**

1、管理员登陆

```bash
db.auth("root","123456");
```

2、创建数据库或切换数据库

```bash
use demo
```

3、创建其他用户并授权访问demo数据库

```
db.createUser({user:'test',pwd:'123456',roles:[{role:"readWrite",db:"demo"}]});
```

# 数据库相关操作

**一、创建和切换数据库**

* 使用`use`创建和切换数据库
* 数据库只有在内容插入后才会持久化到磁盘

```bash
use articledb;
```

**二、查看所有数据库**

* 使用`show dbs`查看所有有权限查看的数据库

```bash
show dbs;
```

**三、查看当前数据库**

* 使用`db`查看当前数据库

```bash
db;
```

**四、删除数据库**

* 用于删除持久化后的数据库
* 使用`db.dropDatabase()`实现

```bash
db.dropDatabase();
```

**五、数据库名的要求**

* 不能是空字符串
* 不能含有特殊字符
* 全部小写
* 最多64字节长度

**六、特殊的数据库**

* **admin**：root数据库，用于权限的管理
* **local**：该数据库不会被复制，存储本地私有数据
* **config**：Mongo用于分片设置时，config数据库在内部使用，用于保存分片信息

# 集合相关操作

* 集合类似于关系型数据库的表

**一、集合的显式创建**

* 使用`db.createCollection()`创建集合

```bash
db.createCollection("mycollection");
```

**二、集合的隐式创建**

* 向集合插入文档时，如果集合不存在则自动创建集合

**三、集合的删除**

* 使用`db.集合名.drop()`删除集合

```bash
db.mycollection.drop();
```

**四、查看库中的表**

* 使用`show collections`或`show tables`查看表

```bash
show collections;
```

# 文档的增删改操作

## 文档的插入

**一、单文档插入**

* 使用`db.集合.insert()`插入一条数据
* 值用json形式的键值对表示
* 插入文档时会实现隐式创建集合

```bash
//插入语法
db.collection.insert(
<document or array of documents>,
{
    writeConcern: <document>,    //可选字段
    ordered: <boolean>    //可选字段
    }
)
```

| 参数                              | 作用                                                         |
| --------------------------------- | ------------------------------------------------------------ |
| \<document or array of documents> | 可设置插入一条或多条文档。                                   |
| writeConcern:\<document>          | 自定义写出错的级别                                           |
| ordered:\<boolean>                | 如果为 true，在数组中执行文档的有序插入，发生错误则返回未处理的文档<br>如果为 false，则执行无序插入，若其中一个文档发生错误，则忽略错误 |

```bash
//向comment集合插入文档
db.comment.insert({"artivleid":"1","content":"hello world"})
```

**二、多文档插入**

* 使用`db.文档.insertMany()`插入多条数据
* 多条数据用数组形式包裹起来

```bash
db.comment.insertMany([{"articleid":"2","content":"two"},{"articleid":"3","content":"three"}])
```

**三、捕获错误**

* 使用`try{ 插入语法 }catch(e){print(e);}`捕获错误
* 如果某条数据插入失败将会终止插入，成功插入的数据不会回滚

```bash
try{
	db.comment.insertMany([
{"articleid":"2","content":"two"},{"articleid":"3","content":"three"}
]);
}catch(e){
	print(e);
}
```

## 文档的更新

**一、覆盖修改**

* 使用`db.集合.update()`实现覆盖修改，如果只其中某些值，则其它值会删除
* 语法：`db.集合.update({定位键值对},{修改的键:})`
* 输入数字时，需要使用`键:NumberInt()`，否则默认为浮点数
* 默认修改第一条数据

**二、局部修改**

* 使用`db.集合.update()`实现局部修改
* 使用`$set`修改器修改键值对的值
* 语法：`db.集合.update({定位键值对},$set:{键:值})`
* 默认修改第一条数据

**三、修改多条数据**

* 修改语句后面追加`,{multi:true}`实现修改多条数据

**四、列值的自动增长**

* 使用`$inc:{}`实现列值自增
* 语法：`db.集合.update({定位键值对},{$inc:{键:NumberInt(自增数字)}})`

## 文档的删除

**一、删除文档**

* 使用`db.集合.remove(条件)`删除文档

**二、清空集合数据**

* 使用`db.集合.remove({})`清空集合数据

```bash
db.comment.remove({})
```

**三、使用_id删除数据**

```bash
db.comment.remove({_id:"1"})
```

# 文档的查询操作

## 基本查询

**一、查询集合的所有文档**

* 使用`db.集合.find()`查询集合所有文档

```bash
db.comment.find()
```

**二、根据某个键值对查询所有符合的文档**

* 使用`db.集合.find({key:"value"})`查询
* key不用引号

```bash
db.comment.find({articleid:"2"})
```

**三、根据某个键值对查询一条符合的文档**

* 使用`db.集合.findOne({key:"value"})`查询
* 返回第一条查询的数据

```bash
db.comment.findOne({articleid:"2"})
```

**四、选择字段查询**

* 使用`db.集合.find({key:"value"},{key:bool,key:bool})`查询
* bool值用0和1表示

## 统计查询

**一、统计所有记录**

```bash
db.commemt.count()
```

**二、条件统计**

* 格式：`db.集合名.count({条件字段:"值"})`

```bash
db.comment.count({key:"value"})
```

## 分页查询

**一、查询指定数量的记录**

* 使用`.limit(数量)`实现

```bash
db.comment.find().limit(3)
```

**二、分页查询**

* 使用`skip(数量)`跳过前几条记录，使用`limit(数量)`实现查询指定数量记录

```bash
db.comment.find().skip(2).limit(2)
```

## 排序查询

**一、排序查询**

* 使用`sort()`指定参数进行排序查询，值为1时升序，值为-1时降序
* 输入多个键值对时，越前的键值对排序优先级越大

```bash
#按uid升序查询
db.comment.find().sort({uid:1})
#按uid降序查询
db.comment.find().sort({uid:-1})
```

## 正则查询

**一、正则查询**

* 输入值为正则表达式的键值对升序正则查询

```bash
db.comment.find({content:/^你好/})
```

## 比较查询

**一、比较查询**

* 格式：`db.集合名.find({"字段":{比较操作符:值}})`

```bahs
db.comment.find({like:{$gt:NumberInt(100)}})
```

| 比较操作符 | 描述     |
| ---------- | -------- |
| $gt        | 大于     |
| $lt        | 小于     |
| $lte       | 小于等于 |
| $ne        | 不等于   |

## 逻辑查询

**一、与或查询**

* 使用`$and`或`$or`实现
* 格式：`db.集合.find({逻辑操作符:[{条件1},{条件2}]})`

```bash
db.comment.find({$or:[{a:{$gt:NumberInt(100)}},{b:{$gt:NumberInt(100)}}]})
```

# 索引

## 简介

**一、索引**

* mongodb使用bTree索引

## 索引类型

**一、类型**

* 单字段索引：给某个字段加索引
* 多字段索引：给多个字段加索引，和mysql一样

## 索引相关操作

**一、查看索引**

* 使用`db.集合.getIndexes()`获取集合的所有索引

```bash
db.coooment.getIndexes()
```

**二、创建索引**

* 格式：`db.集合.createIndex(keys,options)`

```bash
#单字段升序索引
db.comment.createIndex({uid:1})
#多字段索引
db.comment.createIndex({uid:1},{name:-1})
```

| 可选项 | 描述                       |
| ------ | -------------------------- |
| unique | 唯一索引，要求字段不可重复 |
| name   | 索引名                     |

**三、删除索引**

* 格式：`db.集合.dropIndex(索引名)`

**四、删除所有索引**

* 格式：`db.集合.dropIndexs()`

# Go连接mongo

## 驱动下载及相关文档

**一、驱动下载**

```bash
go get go.mongodb.org/mongo-driver/mongo
```

**二、相关文档**
[The Official Golang driver for MongoDB (github.com)](https://github.com/mongodb/mongo-go-driver)

## 连接数据库

**一、获取连接**

* 使用`options.Client()`获取配置对象
* 使用`.ApplyURI()`设置连接地址(`mongodb://IP:端口`)
* 使用`.SetMaxPoolSize()`设置连接池大小
* 使用`.SetAuth()`设置用户和密码，填入`options.Credential{}`结构体（可选）

* 使用`mongo.Connect()`连接
* 使用`.Ping()`测试连接

```go
func conn() (conn *mongo.Client, err error) {
	//定义参数
	option := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	//可设置连接池
	option.SetMaxPoolSize(10)
    //设置用户和密码
	option.SetAuth(options.Credential{Username: "root", Password: "123456"})
	//获取连接
	conn, err = mongo.Connect(context.Background(), option)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//测试连接
	err = conn.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connect success")
	return
}
```

**二、连接数据库**

* 获取`*mongo.Client`对象后进行连接
* 使用`.Database()`连接到集合，使用`.Collecion()`连接到文档

```go
client:=conn.Database("test").Collection("article")
```

## 插入数据

**一、单条插入**

* 使用结构体定义字段，然后插入数据
* 使用`.InsertOne()`插入一条数据
* 返回`_id`序列

```go
type demo struct{
    Id	int64	`json:"id"`
    Text	string	`json:"text"`
}

d:=&demo{Id:1,Text:"hello"}
res,err:=client.InsertOne(context.Background(),d)
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(res.InsertedID)
```

**二、多条插入**

* 使用`.InsertMany()`插入多条数据

```go
d1:=&demo{Id:1,Text:"hello"}
d2:=&demo{Id:2,Text:"world"}
var list []interface{}
list=append(list,d1,d2)
res,err:=client.InsertMany(context.Background(),list)
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(res)
```

## 查询数据

**一、单条件查询和单条查询**

* 使用`bson.M{}`定义单个条件
* 值为数字时使用数字，不要使用字符串

* 使用`.FindOne()`查询一条数据
* 使用`.Decode()`解码到结构体

```go
where:=bson.M{"id":4}
var res demo
err=client.FindOne(context.Background(),where).Decode(&res)
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(res)
```

**二、_id查询**

* 需要使用`primitive.ObjectIDFromHex()`将id序列转为`ObjectID`类型
* 使用`.Decode()`解码到结构体

```go
objID,_:=primitive.ObjectIDFromHex("64093e8eba5e56b5e3277c81")
where:=bson.M{"id":objID}
var res demo
err=client.FindOne(context.Background(),where).Decode(&res)
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(res)
```

**三、多条查询**

* 使用`.Find()`获取多条数据并返回游标
* 通过游标获取值解码到结构体
* 使用`.Decode()`解码到结构体

```go
var res []*demo
//获取游标
cur,err:=client.Find(context.Background(),bson.D{{}})
if err!=nil{
    fmt.Println(err)
    return
}
//游标取值
for cur.Next(context.Background()){
    var d demo
    cur.Decode(&d)
    res=append(res,&d)
}
for _,v:=range res{
    fmt.Println(v)
}
```

**四、分页查询**

* 获取查询配置对象`&options.FindOptions{}`
* 使用`.SetLimit()`设置限制参数
* 使用`.SetSkip()`设置跳过参数

```go
//设置查询选项
findOptions:=&options.FindOptions{}
findOption.SetLimit(3)
findoption.SetSkip(1)
var res []*demo
//放入查询选项获取游标
cur,err:=client.Find(context.Background(),bson.D{{}},findOptions)
if err!=nil{
    fmt.Println(err)
    return
}
//游标取值
for cur.Next(context.Background()){
    var d demo
    cur.Decode(&d)
    res=append(res,&d)
}
for _,v:=range res{
    fmt.Println(v)
}
```

## 更新数据

**一、更新数据**

* 使用`bson.D{}`用`$set`设置局部更新（不设置会导致这个文档的部分字段归零）

* 使用`.UpdateOne()`更新一条文档

```go
where:=bson.M{"id":4}
update:=bson.D{{"$set",bson{{"text,"update data}}}}
res,err:=client.UpdateOne(context.Background(),where,update)
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(res.MatchedCount)
```

## 删除数据

**一、删除数据**

* 使用`.DeleteOne()`删除一条数据
* 使用`.DeleteMany()`删除多条数据
* 使用`.Drop()`删除所有数据





