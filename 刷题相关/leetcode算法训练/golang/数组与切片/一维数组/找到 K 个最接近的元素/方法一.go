package main

import (
	"fmt"
	"sort"
)

func do(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool {				//使用sort包的SliceStable接口设计排序算法
		if cal(arr[i], x) < cal(arr[j], x) {		//SliceStable函数里比较式子遇到相等情况时会保留原始顺序，不会对顺序进行处理
			return true
		}
		return false
	})
	arr = arr[:k]											//截取k个元素
	sort.Ints(arr)											//重新排序
	return arr												//返回
}

func cal(n int, x int) int {								//计算并返回绝对值
	if n-x < 0 {
		return x - n
	}
	return n - x
}

func main() {
	arr := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}				
	//arr := []int{1, 2, 3, 4, 5}
	fmt.Println(do(arr, 3, 5))								//处理耗时长
}
