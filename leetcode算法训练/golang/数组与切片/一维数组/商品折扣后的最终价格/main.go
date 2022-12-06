package main

import "fmt"

func do(prices []int) []int {
	var discount int								//定义折扣

	for index, val := range prices {				//遍历商品价格
		discount = 0								//默认折扣为0
		for i := index + 1; i < len(prices); i++ {	//索引初始值为下一个商品
			if prices[i] <= val {					//如果后面商品的价格存在小于等于当前商品的价格
				discount = prices[i]				//确定折扣
				break
			}
		}
		prices[index] = val - discount				//根据折扣价修改商品价格
	}
	return prices
}

func main() {
	arr := []int{8, 4, 6, 2, 3}
	fmt.Println(do(arr))
}
