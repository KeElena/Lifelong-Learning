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
	dt := time.Now().Add(time.Second * 5)
	ctx, cancel := context.WithDeadline(context.Background(), dt)
	defer cancel()

	go do(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("5s after")
	time.Sleep(time.Second * 3)
}
