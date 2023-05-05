package main

import (
	"fmt"
	"strings"
)

func do(words []string) (res []string) {
	
	for i, childStr := range words {					//遍历子串
		for j, fathStr := range words {					//遍历父串
			if i == j {									//索引相同则跳过
				continue
			}
			if strings.Contains(fathStr, childStr) {	//父串包含子串时追加
				res = append(res, childStr)
				break									//去冗余
			}
		}
	}
	return												//返回
}

func main() {
	arr := []string{"leetcoder", "leetcode", "od", "hamlet", "am"}
	fmt.Println(do(arr))
}
