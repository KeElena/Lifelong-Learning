package main

import (
	"fmt"
	"strconv"
	"strings"
)

func do(equation string) string {

	arr := make([]int, 2, 2)				//arr[0]表示x，arr[1]表示b，组成0=ax+b
	i, opp := 2, -1							//必要参数，i为索引，opp为表示相反数的参数
	var x, b int
	if equation[0] != '-' {					//增加前缀
		equation = "+" + equation
	}
	equation = equation + "+"				//增加后缀

	for len(equation) > 1 {					//equation="+"时退出循环

		if equation[0] == '=' {				//左右式子分界处
			opp = 1							//相反参数变为1
			equation = equation[1:]			//剔除=
			if equation[0] != '-' {			//增加前缀
				equation = "+" + equation
			}
			continue						//跳过循环
		}

		for {								//获取下一个+或-或=的索引位置（不是第0个）
			if (equation[i] == '+' || equation[i] == '-' || equation[i] == '=') && len(equation) >= 2 {
				break
			}
			i++
		}
											//包含x则执行
		if strings.Contains(equation[:i], "x") {
			if len(equation[0:i]) == 2 {	//根据长度判断有无参数
				x = 1
			} else {
				x, _ = strconv.Atoi(equation[1 : i-1])
			}
		} else {							//没有x则提取常数
			b, _ = strconv.Atoi(equation[1:i])
		}
										
		if equation[0] == '+' {				//执行+法相关运算
			arr[0] += x * opp				//每轮循环会重置x和b为0
			arr[1] += b * opp
		} else {							//执行-法运算
			arr[0] -= x * opp
			arr[1] -= b * opp
		}
		equation = equation[i:]				//剔除处理过的数据
		x = 0								//参数重置
		b = 0
		i = 2
	}
	if arr[0] == 0 && arr[1] == 0 {			//判断因为明确解
		return fmt.Sprintf("Infinite solutions")
	}
	if arr[0] == 0 && arr[1] != 0 {			//判断无解
		return fmt.Sprintf("No solution")
	}
	return fmt.Sprintf("x=%d", -arr[1]/arr[0])//返回解
}

func main() {
	str := "-x=-1"
	fmt.Println(do(str))
}
