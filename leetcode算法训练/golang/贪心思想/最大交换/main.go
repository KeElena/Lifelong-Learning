package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {

	Byte := []byte(strconv.Itoa(num))						//处理字符串的内容时最好转为字节切片
	var max byte
	var idx int
	var ok bool

	for i := 0; i < len(Byte); i++ {
		if !ok && i < len(Byte)-1 && Byte[i] >= Byte[i+1] {	//跳过开头递减的情况
			continue
		} else {
			ok = true
		}

		if Byte[i] >= max {									//大于等于max则保存值和索引，要求值最大，且相等情况下索引要最大
			max = Byte[i]
			idx = i
		}
	}

	for i := 0; i < len(Byte); i++ {						//选第一个比max小的值并交换
		if max > Byte[i] {
			Byte[idx] = Byte[i]
			Byte[i] = max
			break
		}
	}

	num, _ = strconv.Atoi(string(Byte))
	return num
}
//需要跳过开头数值递减的情况，否则开头如果第一个值为最值或前几个值为最大的前几个值，导致交换一次后值永远得不到结果
func main() {
	num := 1993

	fmt.Println(maximumSwap(num))
}
