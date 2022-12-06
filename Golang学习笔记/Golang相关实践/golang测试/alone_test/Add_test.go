package test_one

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	want := 15
	result := Add(arr)
	if reflect.DeepEqual(result, want) == false {
		t.Fatalf("arr:%v want=%d result=%d failed!", arr, want, result)
	} else {
		fmt.Println("OK")
	}
}
