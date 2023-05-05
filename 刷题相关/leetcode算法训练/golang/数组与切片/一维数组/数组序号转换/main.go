package main

import (
	"fmt"
	"sort"
)

func do(arr []int) []int {
	mapArr, cp, res, num := make(map[int]int, len(arr)), append(make([]int, 0, len(arr)), arr...), make([]int, 0, len(arr)), 1

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if i-1 > -1 && arr[i-1] == arr[i] {
			continue
		}
		mapArr[arr[i]] = num
		num++
	}

	for i := 0; i < len(cp); i++ {
		res = append(res, mapArr[cp[i]])
	}
	return res
}

func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(do(arr))

}
