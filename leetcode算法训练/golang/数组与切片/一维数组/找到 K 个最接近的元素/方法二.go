package main

import (
	"fmt"
	"sort"
)

func do(arr []int, k int, x int) []int {

	res := make([]int, 0, k)					//初始化结果集
	m := 0										//定义一个指针
	//获取最接近x元素的索引
	for {
		if m+1 == len(arr) {					//指针位置在最右边时退出循环
			break
		}
		if cal(arr[m], x) < cal(arr[m+1], x) {	//下一个元素与x的差值比当前元素大时返回
			break
		}
		m++										//指针+1
	}
	//定义左右指针
	n := m										//定义右指针
	m--											//定义左指针
	//循环k次获取k个值
	for i := 0; i < k; i++ {
		//左指针的差值小于等于右指针的差值则追加左指针的值
		if n < len(arr) && m >= 0 && cal(arr[m], x) <= cal(arr[n], x) {	//题目要求数小优先
			res = append(res, arr[m])
			m--															//左指针-1
			continue
		}
		//右指针的差值小于左指针的差值则追加右指针的差值
		if n < len(arr) && m >= 0 && cal(arr[m], x) > cal(arr[n], x) {
			res = append(res, arr[n])
			n++															//右指针+1
			continue
		}
		//当左指针超出数组长度范围时对右指针的元素进行追加
		if m < 0 {
			res = append(res, arr[n])
			n++															//右指针+1
			continue
		}
		//当右指针超出数组长度范围时对左指针的元素进行追加
		if n >= len(arr) {
			res = append(res, arr[m])
			m--															//左指针-1
			continue
		}
	}
	//对结果进行排序
	sort.Ints(res)
	return res
}

func cal(n int, x int) int {
	if n-x < 0 {
		return x - n
	}
	return n - x
}

func main() {
	//arr := []int{0, 0, 1, 2, 3, 3, 4, 4, 7, 7, 8}
	//arr := []int{1, 2, 3, 4, 5}
	arr := []int{1, 3}
	fmt.Println(do(arr, 1, 2))											//使用到的排序算法少，处理速度快
}
