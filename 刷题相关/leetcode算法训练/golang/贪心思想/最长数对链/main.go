package main

import (
	"fmt"
	"sort"
)

func do(pairs [][]int) int {

	var lastIndex, res int					//定义需要使用的变量

	sort.Slice(pairs, func(i, j int) bool {	//对二维数组进行排序
		if pairs[i][1] < pairs[j][1] {		//下一个数组的第二个元素小于当前数组的第二个元素时，进行位置交换
			return true
		}
		return false
	})
											//lastIndex默认为0
	res = 1									//设置数对链的最低结果
	for i := 1; i < len(pairs); i++ {		//一位lastIndex默认为0，所以从第2个元素开始循环遍历
		if pairs[i][0] > pairs[lastIndex][1] {
			lastIndex = i					//每次循环用当前索引数组的第1个元素与lastIndex位置对应数组的第2个值进行比较
			res++							//大于则可以组成数对链，将当前索引作为lastIndex参数，用于判断后面元素能否组成数对链
		}

	}
	return res
}
//题目数对的特点：第一个元素永远比第二个元素小
//数对链的特点：前一个数对的第2个元素比后一个数对的第1个元素小
//数对的第一个元素决定能否组成数对链
//数对的第二个元素能影响组成数对链的数对个数
//要获取最长的数对链，需要对二维数组进行排序，根据每个数对的第二个元素的大小从小到大进行排序，然后判断前后数对能否组成数对链
func main() {
	//arr := [][]int{{1, 2}, {7, 8}, {4, 5}}
	arr := [][]int{{7, 9}, {4, 5}, {7, 9}, {-7, -1}, {0, 10}, {3, 10}, {3, 6}, {2, 3}}
	fmt.Println(do(arr))
}
