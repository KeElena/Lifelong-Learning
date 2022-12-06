package main

import "fmt"

func do(s string) int {

	var sum, maxScore int
	for i := 1; i <= len(s)-1; i++ {		//左子串可能长度为1到len(s)-1，控制左字串可以控制右子串长度
		for j, char := range s {
			if j < i {
				if char == '0' {			//第i个元素为右子串元素，左边的全为左子串元素
					sum++
				}
			} else {
				if char == '1' {
					sum++
				}
			}
		}
		if sum > maxScore {					//保存最大值
			maxScore = sum
		}
		sum -= sum							//每次内循环结束重置1次
	}
	return maxScore
}

func main() {
	str := "00"
	fmt.Println(do(str))
}
