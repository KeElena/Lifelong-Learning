# ORM

* 全称：对象关系映射

* 映射关系：

| 数据库 | 结构体     |
| ------ | ---------- |
| 数据表 | 结构体     |
| 数据行 | 结构体实例 |
| 字段   | 结构体字段 |

* ORM的优缺点

| 优点 | 提高开发效率<br>安全                      |
| ---- | ----------------------------------------- |
| 缺点 | 牺牲执行性能<br>牺牲灵活性<br>弱化SQL能力 |

# Gorm

**一、文档**

* [Gorm中文文档](https://gorm.io/zh_CN/docs/)

**二、下载**

```shell
go get -u gorm.io/gorm
#mysql驱动
go get -u gorm.io/driver/mysql
```

# 数据库连接

**一、需要的参数**

* 连接名：个人使用一般为root
* 密码
* ip
* 端口：一般为3306
* 数据库名
* 字符编码：最好是utf8mb4

```go
var (
	user	="root"
    pwd		="123456"
    ip		="127.0.0.1"
    port	="3306"
    db		="demo"
    charset	="utf8mb4"
)
```

**二、获取数据库连接**

* 使用`gorm`官方提供的`mysql`驱动解析数据源
* 使用`gorm.Open()`获取数据源，包含多个数据库连接

```go
var (
	user    = "root"
	pwd     = "123456"
	ip      = "127.0.0.1"
	port    = "3306"
	db      = "demo"
	charset = "utf8mb4"
)
//获取数据源
dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pwd, ip, port, db, charset)
//获取数据源，需要高级设置可以修改config结构体
engine, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
if err != nil {
	fmt.Println(err)
	return
}
```

# 数据插入

## 单行插入

* 使用`create()`插入数据
* 使用`Table()`指定表

```go
var err error
//编写数据
value:=&User{Name:"zhangsan",Age:22,Birthday:"2000-01-01"}
//指定表插入数据
err=engine.Table("user").Create(value).Error
//获取错误，判断是否插入成功
if err!=nil{
    fmt.Println(err)
    return
}
```

## 批量插入

* 使用`CreateInBathes()`进行批量插入数据
* 填入的数字表示一次上传n个数据，越大可缩短SQL语句减少访问数据库的次数

```go
var ValueSlice []*User
var num int64
value1:=&User{Name:"zhangsan",Age:22}
value2:=&User{Name:"lisi",Age:22}
ValueSlice=append(ValueSlice,value1,value2)
num=engine.Table("user").CreateInBatches(ValueSlice,3).RowsAffected
if num!=len(ValueSlice){
    fmt.Println("数据插入错误或数据插入不完整!")
}
```

# 数据更新

## 强制更新行的所有字段数据

* 使用`Save()`更新数据
* `Save()`方法更新数据时，如果某些字段没有输入值，则会把对应字段的数据置为零值

```go
var err error
value1:=&User{Name:"zhangsan",Birthday:"2000-02-02"}
err=engine.Table("user").Where("id=?",19).Save(&value1).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println("ok")
```

## 选择字段更新数据

**一、单字段数据更新**

* 使用`Update()`实现单字段数据更新
* 使用`Model()`指定表才能触发`hook`函数，`Table()`不会触发

```go
err=engine.Table("user").Where("id=?",19).Update("age",19).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println("ok")
```

**二、多字段数据更新**

* 使用`Updates()`实现多字段数据更新
* 使用`map`或`struct`填充数据
* 使用`Model()`指定表才能触发`hook`函数，`Table()`不会触发

```go
value:=make(map[string]interface{},2)
value["age"]=12
value["name"]="zhangqi"
err=engine.Table("user").Where("id=?",20).Updates(&value).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println("ok")
```

## 不使用hook更新数据

* 当结构体有`gorm.Model`，且表里有`updated_at`字段时，更新数据会自动更新`updated_at`的时间
* 使用`UpdateColumn()`更新单字段数据不会触发hook函数
* 使用`UpdateColumns()`更新多字段数据不会触发hook函数
* `hook`函数最好在数据库里设置触发器实现而不是在代码通过`hook`实现

## 选择更新部分字段

* 使用`select()`可以选择字段进行更新
* 使用`Omit()`可以剔除某些字段进行更新

# 删除数据

* 使用`Delete()`删除数据
* 使用结构体条件选择数据

```go
err=engine.Table("user").Delete(&User{Id:19}).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println("ok")
```

# 查询数据

## 查询一条数据

**一、查询一条数据**

* 使用`Take()`查询1条数据

```go
value:=User{}
err=engine.Table("user").Where("age>?",19).Take(&value).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(value)
//相当于sql语句：select * from user where age>19 limit 1;
```

**二、查询主键正序第一条数据**

* 使用`First()`获取数据
* 获取的是主键正序第一条数据

```go
value:=User{}
err=engine.Table("user").First(&value).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(value)
//相当于sql语句：select * from user order by id asc limit 1;
```

**三、查询最后一条数据**

* 使用`Last()`获取数据
* 获取的是主键倒序第一条数据

```go
value:=User{}
err=engine.Table("user").Last(&value).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(value)
//相当于sql语句：select * from user order by id desc limit 1;
```

## 查询多条数据

* 使用`Find()`查询多条数据

```go
var ValueSlice []User
err=engine.Table("user").Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 单字段查询

* 使用`Pluck()`选择字段进行查询

```go
var ValueSlice []int
err=engine.Table("user").Pluck("age",&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 多字段查询

* 使用`Select()`筛选字段
* 使用`Find()`查询多条数据
* `select()`里可以使用聚合函数

```go
var ValueSlice []User
err=engine.Table("user").Select("name","age").Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 条件查询

* 使用`Where()`进行条件查询

```go
var ValueSlice []User
err=engine.Table("user").Where("age in (?)",[]int{19,20}).Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 排序查询

* 使用`Order()`选择字段和排序方式进行排序
* 输入的值至少要含有两个用空格隔开的值，一个为字段一个为排序规则

```go
var ValueSlice []int
err=engine.Table("user").Select("age").Order("age desc").Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 分页查询

* 使用`Linit()`限制查询数量
* 使用`Offset()`设置起点
* (n-1)*limit

```go
var ValueSlice []User
err=engine.Table("user").Limit(2).Offset(3).Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 分组查询

* 在`Select()`里可以写聚集函数，使用聚集函数产生的新列需要起别名才能被后续调用
* 使用`Having()`进行分组条件选择
* `Group()`必须搭配`Select()`一起使用，因为经常使用聚集函数
* 需要构建一个新的结构体，要求字段名和表的字段和sql执行过程中的别名对应，且首字母大写

```go
type result struct{
    Birthday string
    Num int
}
func main(){
    ...
    数据库连接
    ...
    
    var ValueSlice []result
    err=engine.Table("user").Select("birthday,count(*) as num").Group("birthday").
    Having("num>1").Find(&ValueSlice).Error
    if err!=nil{
        fmt.Println(err)
        return
    }
}
```

# 会话

## 会话的作用

* 避免公用一个数据库连接导致的问题，将db操作分离互不影响

## 会话相关配置

```go
type Session struct{
    DryRun						bool	//生成sql语句但不执行
    PrepareStmt					bool	//预编译模式
    NewDB 						bool	//新的数据库连接，不带之前的条件
    Initialized					bool	//初始化数据库连接，不保证协程安全
    SkipHooks					bool	//跳过钩子
    SkipDefaultTransaction		bool	//禁用默认事务
    DisableNestedTransaction 	bool	//禁用嵌套事务
    AllowGlobalUpdate			bool	//允许不带条件更新
    FullSaveAssociations		bool	//允许更新关联数据
    QueryFields					bool	//select字段
    Context						context.Context
    Logger						logger.Interface
    NowFunc	func()				time.Time//改变GORM获取当前时间的实现
    CreateBatchSize				int		//批量处理大小
}
```

## 获取会话对象与使用

* 使用`Session()`获取会话对象
* 需要输入`gorm.Session{}`结构体

```go
session:=engine.Session(&gorm.Session{})
var ValueSlice []User
err=session.Table("user").Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

## 禁用默认事务

* 在简单查询里禁用事务能够提高性能，默认是开启事务
* 将会话参数`SkipDefaultTransaction`设为`true`

```go
session:=engine.Session(&gorm.Session{SkipDefaultTransaction:true})
var ValueSlice []User
err=session.Table("user").Find(&ValueSlice).Error
if err!=nil{
    fmt.Println(err)
    return
}
fmt.Println(ValueSlice)
```

# 事务

## 自动事务

* 使用`Transaction()`实现自动事务
* 需要实现匿名方法，如果返回的`error`非空，则自动回滚事务
* 使用`tx *gorm.DB`执行操作

```go
err=engine.Transaction(func(tx *gorm.DB)error{
    //更新数据
    err=tx.Table("user").Where("id=?",11).Update("age",17).Error
    if err!=nil{
        return err
    }
    //return fmt.Errorf("error")	//用于主动返回错误
    //插入数据
    err=tx.Table("user").Create(&User{Name:"zhangsan",Age:20}).Error
    if err!=nil{
        return err
    }
    //返回空时提交事务
    return nil
})
```

## 嵌套事务

* 在自动事务的基础上开启新的自动事务
* 子事务如果报错，只要没有接收`error`则不会影响父事务和其他子事务的提交

```go
err=engine.Transaction(func(tx *gorm.DB)error{
    //父事务相关操作
    ...
    //子事务报错不会导致父事务回滚
    engine.Transaction(func(tx2 *gorm.DB)error{
    	//子事务相关操作
    	return nil
	})
    //子事务报错不会导致父事务回滚
    engine.Transaction(func(tx2 *gorm.DB)error{
    	//子事务相关操作
    	return nil
	})
    return nil
})
```

## 手动事务

* 使用`Begin()`获取事务对象，并开启事务
* 使用`tx.Rollback()`回滚事务
* 使用`tx.Commit()`提交事务

```go
//开启事务，获取事务对象
tx:=engine.Begin()
//更新数据
err=tx.Table("user").Where("id=?",25).Update("age",18).Error
if err!=nil{
    fmt.Println(err)
    tx.Rollback()
}
//插入数据
err=tx.Table("user").Create(&User{Name:"wangliu",Age:18,Birthday:"1998-01-01"}).Error
if err!=nil{
    fmt.Println(err)
    tx.Rollback()
}
//提交数据
tx.Commit()
```

## 事务保存点

* 事务回滚时可以回滚到某个保存点，而不用将所有数据都回滚

* 使用`tx.SavePoint()`设置保存点
* 使用`tx.RollbackTo()`回滚到某个保存点

```go
//设置保存点
tx.SavePoint("sp1")
//回滚到保存点
tx.RollbackTo("sp1")
```

# hook钩子函数

**一、hook**

* Hook在插入、更新、删除、查找等操作之前或之后调用的函数
* <font color=red>如果hook函数返回了错误，则会造成事务的回滚</font>

* 如果事务发生回滚，则hook造成的影响也会回滚
* hook函数类型：`func(*gorm.DB)`
* 使用`hook`函数需要用`Model()`或在操作函数里输入结构体才能触发

**二、创建hook函数**

* `BeforeSave()`：保存更新前，所有更新数据的操作都会触发（增删改）
* `BeforeCreate()`：插入前，只有插入才能触发
* `AfterCreate()`：插入后，只有插入才能触发
* `AfterSave()`：保存更新后，所有更新数据的操作都会触发（增删改）
* `BeforeFind()`：查询了n条会触发n次，没有结果不会触发

```go
//创建钩子函数
func (user *User)BeforeCreate(db *gorm.DB)error{
    fmt.Println("insert before")
    return nil
}
func (user *User)BeforeSave(db *gorm.DB)error{
    fmt.Println("data update before")
    return nil
}
```

```go
//调用钩子函数
tx := engine.Begin()
//更新数据时将结构体放入Model()
err = tx.Model(&User{}).Where("id=1").Update("age", 1).Error
if err != nil {
	fmt.Println(err)
	tx.Rollback()
	return
}
//插入数据时在Create()里放入结构体
err = tx.Table("user").Create(&User{Name: "wangliu", Age: 18, Birthday: "21998-01-01"}).Error
if err != nil {
	fmt.Println(err)
	tx.Rollback()
	return
}
tx.Commit()
```

# 常见方法汇总

| 方法              | 作用                                                         |
| ----------------- | ------------------------------------------------------------ |
| Create()          | 插入数据                                                     |
| CreateInBatches() | 批量插入数据                                                 |
| Table()           | 指定表                                                       |
| RowsAffected      | 获取受影响行数                                               |
| Error             | 获取错误                                                     |
| Exec()            | 执行原生sql语句，使用？作为占位符，不能获取值                |
| Raw()             | 支持使用占位符？绑定多参数，执行原生sql语句并取值            |
| Model()           | 放入结构体，可用于指定表，<font color=red>默认在结构体名的小写后面加s</font><br>用于指定结构体，以便可以用对应结构体的hook函数<br>优先级比Table()低，如果有Table()则用Table()对应的表 |
| Select()          | 选择字段，可以使用聚合函数                                   |
| Omit()            | 剔除字段                                                     |
| Take()            | 单行查询                                                     |
| First()           | 获取第一条数据                                               |
| Last()            | 获取最后一条数据                                             |
| Find()            | 获取多行数据                                                 |
| Pluck()           | 单字段多行查询                                               |
| Scan()            | 多行查询，但要求必须用Model()和Table()指明表                 |
| Transaction()     | 开启自动事务                                                 |
| Begin()           | 获取事务对象并开启事务                                       |
| tx.Rollback()     | 回滚                                                         |
| tx.Commit()       | 提交                                                         |
| tx.SavePoint()    | 设置保存点                                                   |
| tx.RollbackTo()   | 回滚到保存点                                                 |







