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
			fmt.Println(ctx.Value("key"))
			time.Sleep(time.Millisecond * 500)
		}
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ctx = context.WithValue(ctx, "key", "value")

	go do(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("after 5s")
	time.Sleep(time.Second * 3)
}
