package main

import "fmt"

func do(groupSizes []int) [][]int {

	res := make([][]int, 0, len(groupSizes)/2)				//结果集
	classMap := make(map[int][]int, len(groupSizes)/2)		//构造哈希表

	for i, val := range groupSizes {						//遍历数组，按出现次数进行归类
		classMap[val] = append(classMap[val], i)
	}

	for key, s := range classMap {							//遍历哈希表
		for len(s) > 0 {
			res = append(res, s[0:key])						//将结果追加到结果集
			s = s[key:]										//剪切数组
		}
	}
	return res
}

func main() {
	arr := []int{2, 1, 3, 3, 3, 2}
	fmt.Println(do(arr))
}
