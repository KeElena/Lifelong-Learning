package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var sum int32

func do(i int) {
	atomic.AddInt32(&sum, 2)
	time.Sleep(time.Second * 2)
	fmt.Println(i)
}

func main() {
	defer ants.Release()

	pool, _ := ants.NewPool(20)
	defer pool.Release()

	task := func(i int) func(){
		return func() {
			do(i)
			wg.Done()
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		_ = pool.Submit(task(i))
	}
	wg.Wait()

	fmt.Println("sum:", sum)
}