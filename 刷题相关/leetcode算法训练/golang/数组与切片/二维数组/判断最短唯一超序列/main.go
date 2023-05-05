package main

import "fmt"

func do(nums []int, sequences [][]int) bool {

	length := len(nums)
	var Map = make(map[int][]int, length-1)                         //使用map切片混合结构存储“有向图”

	for _, arr := range sequences {                                 //遍历二维数组获取一维数组
		for i := range arr {                                        //遍历一维数组
			if i == 0 {                                             //跳过第一个循环
				continue
			}
			bol := true
			for _, v := range Map[arr[i-1]] {                       //遍历map的切片元素判断是否有重复路径
				if v == arr[i] {
					bol = false                                     //有则置为false
					break                                           //跳出循环
				}
			}
			if bol {
				Map[arr[i-1]] = append(Map[arr[i-1]], arr[i])       //给map的切片追加元素
			}
		}
	}

	sum := 0
	for i := 0; i < length; i++ {                                   //循环遍历nums切片
		if i == 0 {
			continue                                                //跳过第一个循环
		}
		for _, v := range Map[nums[i-1]] {                          //检测map表里是否有该路径
			if v == nums[i] {                                       //有则计数+1
				sum++
			}
		}
	}

	if sum == length-1 {                                            //所有相邻路径都存在说明是唯一的超序列
		return true
	}
	return false
}

func main() {
	var nums = []int{4, 1, 5, 2, 6, 3}
	var sequences = [][]int{
		{5, 2, 6, 3},
		{4, 1, 5, 2},
	}
	fmt.Println(do(nums, sequences))
}
//拓扑排序得到的结果是否唯一取决于结果相邻元素是否有存在一条路径
//将nums转换为拓扑排序的结果
//将sequence二维数组转换为图，相邻元素之间认为有一条路径