package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var sum int32

func do() {

	atomic.AddInt32(&sum, 2)
	time.Sleep(time.Second * 2)
	fmt.Println("hello world")
	fmt.Println("-------")
	wg.Done()
}

func main() {

	t1 := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go do()
	}
	wg.Wait()
	t2 := time.Now()

	fmt.Println("ordinary:", t2.Sub(t1).Seconds())
	fmt.Println("sum:", sum)
}
