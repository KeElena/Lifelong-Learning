package main

import (
	"fmt"
	"testing"
)

func BenchmarkGetUuid(b *testing.B) {
	var count int
	obj, _ := getSnowHost(1, 41, 5)
	for i := 0; i < b.N; i++ {
		_, err := obj.GetUuid()
		if err != nil {
			count++
		}
	}
	fmt.Println(count)
}
