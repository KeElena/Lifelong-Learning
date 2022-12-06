package test_group

import (
	"reflect"
	"testing"
)

type testCase struct {
	arr  []int
	want int
}

func TestAdd(t *testing.T) {
	testGroup := []testCase{
		{arr: []int{1, 2, 3, 4, 5}, want: 15},
		{arr: []int{1, 2, 3, 4, 5, 6}, want: 21},
	}
	for _, test := range testGroup {
		result := Add(test.arr)
		if reflect.DeepEqual(result, test.want) == false {
			t.Fatalf("arr:%v want=%d result=%d failed!", test.arr, test.want, result)
		}
	}
}
