package main

import (
	"fmt"
	"math/rand"
	"time"
)

var(								//声明通道
	chan1 chan int
    chan2 chan int
)
func one(){
    for{
        time.Sleep(time.Second)
        chan1<-rand.Intn(100)
    }
}
func two(){
    for{
        time.Sleep(time.Second)
        chan2<-rand.Intn(100)
    }
}
func main(){
    chan1=make(chan int,3)			//实例化通道
    chan2=make(chan int,3)
    var val int						//声明接收值
    
    go one()						//运行生产者协程
    go two()
    
    for{							//for-select接收产品
        select{						//随机选择通道
            case val=<-chan1:		//chan1有值则进行消费
            	fmt.Println(val)
            case val=<-chan2:		//chan2有值则进行消费
            	fmt.Println(val)
            default:				//通道都没值时休眠
            	fmt.Println("sleep")
            	time.Sleep(time.Second*3)
        }
    }
}