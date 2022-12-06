package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup				//定义全局变量
var sum int32

func do(i int) {					//handle函数
	atomic.AddInt32(&sum, 2)		//原子加法
	time.Sleep(time.Second * 2)
	fmt.Println(i)
}

func main() {
	defer ants.Release()			//程序结束前释放ants

	task := func(i interface{}) {	//使用闭包构建func(i interface{}函数)
		do(i.(int))					//使用断言填充数据
		wg.Done()					//协程计数-1
	}

	p, _ := ants.NewPoolWithFunc(20, task)	//构建协程池，设置协程数和放入工作函数
	defer p.Release()				//协程池程序结束前释放

	for i := 0; i < 20; i++ {		//启动协程
		wg.Add(1)					//协程计数+1
		p.Invoke(i)					//放入数据，至少放入i
	}
	wg.Wait()						//等待运行完毕
	fmt.Println(sum)				//输出结果
}
