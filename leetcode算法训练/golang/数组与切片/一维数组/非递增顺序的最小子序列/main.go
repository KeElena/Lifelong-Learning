package main

import (
	"fmt"
	"sort"
)

func do(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	var sum, temp int
	for _, val := range nums {
		sum += val
	}

	for i, val := range nums {
		sum -= val
		temp += val
		if temp > sum {
			return nums[:i+1]
		}
	}
	return nil
}

func main() {
	nums := []int{7}
	fmt.Println(do(nums))
}
