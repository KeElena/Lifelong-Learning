package main

import "fmt"

func do(distance []int, start int, destination int) int {

	A, B := 0, 0
	for i := start; i < len(distance)*2; i++ {              //顺序做圆周运动
		if i%len(distance) == destination {                 //到底目的则退出循环
			break
		}
		A += distance[i%len(distance)]                      //路径累加，求余计算控制索引范围
	}

	for i := start + len(distance); i > 0; i-- {            //逆向做圆周运动
		if i%len(distance) == destination {                 //到达目的退出
			break
		} 
		B += distance[(i-1)%len(distance)]                 //逆向路径累加
	}

	if A < B {                                             //比较输出最小路径
		return A
	}
	return B
}

func main() {
	arr := []int{3, 6, 7, 2, 9, 10, 7, 16, 11}
	fmt.Println(do(arr, 6, 2))
}
