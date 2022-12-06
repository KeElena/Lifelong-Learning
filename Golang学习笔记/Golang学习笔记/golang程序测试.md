# 单元测试

**一、单元测试的函数类型**

| 类型     | 函数格式        | 作用                       |
| -------- | --------------- | -------------------------- |
| 测试函数 | 前缀为Test      | 测试程序的逻辑行为是否正确 |
| 基准测试 | 前缀为Benchmark | 测试函数性能               |
| 示例函数 | 前缀为Example   | 提供示例文档               |

**二、必要的格式**

​		◼文件名：以`_test.go`为后缀的源代码文件

​		◼必须导入`testing`包

​		◼函数名：以单元测试函数类型的前缀开头（'Test函数名'）

​		◼函数参数：必须要有`t *testing.T`

**三、测试t的作用**

​		◼用于报告测试失败和附加的日志信息

**四、测试一个测试用例**

​		◼需要输入参数、预期值

​		◼运行模块后返回实际值

​		◼使用`reflect.DeepEqual()`方法进行比较，如果false则说明出错

​		◼使用`t.Fatalf()`函数打印出错的测试用例和实际输出值

​		◼测试方式：在终端输入`go test`

```go
//准备测试的模块
func Add(arr []int)int{
    var sum int
    for _,val:=range arr{
        sum += val
    }
    return sum
}
```

```go
//测试函数
func TestAdd(t *testing.T){
    arr :=[]int{1,2,3,4,5}		//输入参数
    want:=15					//期望值
    result:=Add(arr)			//结果值
    if reflect.DeepEqual(result,want)==false{	//预期值与结果值进行比较
        t.Fatalf("arr%v want=%d result=%d failed!")//测试失败则返回测试用例和计算结果
    }else{
        fmt.Println("OK")		//测试成功
    }
}
```

```shell
//开始测试，要求要有go.mod文件
go test
```



**五、测试一组测试用例**

​		◼使用结构体和切片实现测试多组测试用例

```go
//定义测试用例结构体
type testCase struct{
    arr []int
    want int
}
//测试函数
func TestAdd(t *testing.T){
//输入测试用例
    testGroup :=[]testCase{
        {arr: []int{1,2,3,4,5},want:15}		//第一个参数用例
        {arr: []int{1,2,3,4,5},want:21}		//第二个参数用例
    }
//循环测试
    for _,test:=range testGroup{
        result :=Add(test.arr)
        if refect.DeepEqual(result,test.want) ==false{
            t.Fatalf("arr:%v want=%d result %d failed!",test.arr,test.want,result)
        }
    }
}
```

**六、测试一组测试用例且能单独测试某个测试用例**

​		◼使用map对测试用例进行命名

​		◼通过`t.Run()`方法使用匿名函数进行测试

​		◼在终端使用`go test -run=测试函数名/测试用例名`进行单独测试

```go
//定义测试用例结构体
type testCase struct{
    arr []int
    want int
}
//使用map输入测试用例
func TestAdd(t *testing.T){
    testGroup :=map[string]testCase{
    	"case_1":{arr: []int{1,2,3,4,5},want:15},
    	"case_2":{arr: []int{1,2,3,4,5,6},want:21},        
    }
    //遍历map，使用t.Run()方法进行测试
    for name,tc:=range testGroup{
        //使用匿名函数编写测试过程
    	t.Run(name,func(t *testing.T){
        	result :=Add(tc.arr)
        	if reflect.DeepEqual(result,tc.want)==false{
            	t.Fatalf("arr:%v want=%d result=%d failed!",tc.arr,tc.want,result)
        	}
    	})
	}
}
```

```shell
//单独对某个测试用例测试,要求要有go.mod文件
go test -run=TestAdd/case_2
```

**七、测试覆盖率**

​		◼要求被测试的模块和测试文件在同一个包

​		◼被测试模块不能再测试文件内

​		◼有go.mod文件可以直接`go test -cover`好的测试覆盖率

```shell
//获取测试覆盖率，要求要有go.mod文件
go test -cover
```

# 基准测试

**一、基准测试函数格式**

​		◼基准测试以`Benchmark`为前缀，需要一个`*testing.B`参数

​		◼基准测试执行`b.N`次，`b.N`次不固定但一般限制在1秒内执行的次数（`b.N`为常量）

​		◼`func Benchmark函数名(b *testing.B){...}`

**二、计算速度测试**

​		◼使用`for`循环执行模块`b.N`次

​		◼在终端输入指令进行测试：`go test -bench=函数名`

​		◼输出的信息：调用的cpu数量、测试次数、每次执行的速度、测试花费的时间

```go
//斐波那契数列
func Fib(n int)int{
    if n == 1 || n ==2{
        return 1
    }else{
        return Fib(n-1)+Fib(n-2)
    }
}
```

```go
//测试函数
func BenchmarkFib(b *testing.B){
	for i:=0;i<b.N;i++{
		Fib(6)
	}
}
```

```shell
//要有go.mod文件
go test -bench=Fib
```

**三、内存调用情况测试**

​		◼在终端输入指令进行测试：`go test -bench=函数名 -benchmem`

​		◼输出的信息：调用的cpu数、测试次数、每次执行的速度、每次操作消耗的内存、申请内存次数

```shell
go test -bench=Add -benchmen
```

**四、不同参数的性能比较函数**

​		◼先定义一个有输入参数且循环参数的函数

​		◼格式：`func test函数名(b *testing.B,n int){...}`

​		◼再定义一个`func Benchmark函数名 参数值(b *testing.B){...}`的函数

​		◼使用`Benchmark`函数调用`test`函数

​		◼在终端执行`Benchmark`函数进行测试（=后面写函数名+参数）

```go
//定义一个test函数
for testFib(b *testing.B,n int){
    for i:=0i<b.N;i++{
        Fib(n)
    }
}
//定义一个Benchmark函数
func BenchmarkFib10(b *testing.B){
    testFib(b,10)
}
//定义一个不同输入参数的Benchmark函数
func BenchmarkFib20(b *testing.B){
    testFib(20)
}
```

```shell
//执行第一个BenchMark函数
go test -bench=Fib10
//执行第二个BenchMaek函数
go test -bench=Fib20
```

**五、修改一些测试环境**

​		◼`-benchtime=时间s`参数修改测试时间

​		◼`b.ResetTimer`用于重置时间（去掉一些耗时无关操作花费的时间）

​		◼`b.SetParallelism(1)`用于修改使用的CPU数量

```shell
go test -bench=函数名 -benchtime=20s
```

```go
b.ResetTimer()			//重置时间
b.SetParallelism(1)		//测试使用1个核心
```

**六、TestMain测试**

​		◼测试文件如果包含`TestMain(m *testing.M)`，则生成测试时先调用该函数再测试

​		◼`TestMain`再主线程运行，**在调用`m.run`进行测试前或测试后可以进行准备环境或关闭环境的操作**

​		◼`m.run()`内使用匿名方法定义测试函数

# Go性能优化

**一、Go程序性能优化的方向**

​		◼CPU占用情况

​		◼内存占用和使用情况

​		◼死锁情况

​		◼线程使用情况（gorountine profile）

**二、采集性能数据**

​		◼`runtime/pprof`：采集工具型应用运行数据进行分析

​		◼`net/http/pprof`：采集服务型应用运行数据进行分析

​		◼`pprof`开启后每隔一段时间会收集当前堆栈信息，获取所有函数占用的CPU和内存资源，最后做分析

​		◼ 得到数据后使用`go tool pprof`工具进行性能分析

​		◼`go tool pprof`默认使用`-inuse_space`进行统计，还可以用`-inuse-objects`查看分配对象的数量

```go
//工具型应用
import "runtime/pprof"				//导入包
pprof.StartCPUProfile(文件地址)	//开启CPU性能分析
pprof.StopCPUProfile()				//停止CPU性能分析
pprof.WriteHeapProfile(文件地址)	//记录程序的堆栈信息(内存信息)
```

```go
//服务型应用
import _ "net/http/pprof"		//导入包
//导入后HTTP服务会多出/debug/PProf endpoint
```

**三、分析数据**

​		◼采集完数据后在终端输入命令进行分析数据

​		◼shell输入格式：`go tool pprof 数据文件名`（生成数据时文件一般以`.pprof`作为后缀）

​		◼flat：表示当前函数占用CPU耗时（不包含被调用）

​		◼flat%：表示当前函数占用CPU耗时累计的百分比（不包含被调用）

​		◼sum%：函数累计使用CPU的百分比

​		◼cum：当前函数加上调用当前函数的函数占用CPU的总耗时（包含被调用）

​		◼cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比（包含被调用）

​		◼最后一列：函数名称

​		◼进入交互界面后常用的两个指令：`top n`,`list funName`

```shell
go tool pprof cpu.pprof
//进入交互界面
//top n 用于显示占用CPU最多的前几个函数
top 3
//list 函数名 用于获取函数对应的源码内容，用于检查和优化代码
list main.Sum
```

