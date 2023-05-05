package main

import (
	"context"
	"fmt"
	"time"
)

func do(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(1)
			time.Sleep(time.Millisecond * 500)
		}
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go do(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	fmt.Println("end")
	time.Sleep(time.Second * 3)
}
