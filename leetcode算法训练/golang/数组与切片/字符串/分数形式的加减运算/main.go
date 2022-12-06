package main

import (
	"fmt"
)

func do(expression string) string {
	var nume = make([]int, 0, len(expression)/2)		//分子
	var deno = make([]int, 0, len(expression)/2)		//分母
	var commDiv, sum, min = 1, 0, 0						//公倍数
	//分子分母进行分离
	for i := 0; i < len(expression); i++ {
		if expression[i] == '/' {
			if i+2 < len(expression) && expression[i+1] == '1' && expression[i+2] == '0' { //10
				deno = append(deno, 10)
			} else {
				deno = append(deno, int(expression[i+1])-48) //other
			}
			if i-3 > -1 && expression[i-3] == '-' { //-10
				nume = append(nume, -10)
			} else if expression[i-1] == '0' && expression[i-2] == '1' { //10
				nume = append(nume, 10)
			} else if i-2 > -1 && expression[i-2] == '-' { //-other
				nume = append(nume, -(int(expression[i-1]) - 48))
			} else {
				nume = append(nume, int(expression[i-1])-48)
			}
		}
	}
	//计算公倍数
	for _, val := range deno {
		if commDiv%val == 0 {
			continue
		}
		commDiv = commDiv * val
	}
	//计算分子的和
	for i, val := range nume {
		sum += val * commDiv / deno[i]
	}
	//和为0返回
	if sum == 0 {
		return "0/1"
	}

	//取得可能的最小值
	min = sum								//取分子，因为大部分情况下分子的绝对值最小
	if min < 0 {							//判断是否小于0
		min =0								//小于0则取反，以免影响后续循环（由于不是取绝对值的最小值，所以可能不是最小）
	}
	//根据可能最小值min进行约分
	for i := min; i > 1; i-- {
		if sum%i == 0 && commDiv%i == 0 {
			sum /= i
			commDiv /= i
		}
	}
	//返回结果
	return fmt.Sprintf("%d/%d", sum, commDiv)
}

func main() {
	expression := "-1/4-4/5-1/4"
	fmt.Println(do(expression))
}
