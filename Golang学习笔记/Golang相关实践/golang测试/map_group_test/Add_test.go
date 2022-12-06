package test_group_map

import (
	"reflect"
	"testing"
)

type testCase struct {
	arr  []int
	want int
}

func TestAdd(t *testing.T) {

	testGroup := map[string]testCase{
		"case_1": {arr: []int{1, 2, 3, 4, 5}, want: 15},
		"case_2": {arr: []int{1, 2, 3, 4, 5, 6}, want: 21},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			result := Add(tc.arr)
			if reflect.DeepEqual(result, tc.want) == false {
				t.Fatalf("arr:%v want=%d result=%d failed!", tc.arr, tc.want, result)
			}
		})
	}
}
