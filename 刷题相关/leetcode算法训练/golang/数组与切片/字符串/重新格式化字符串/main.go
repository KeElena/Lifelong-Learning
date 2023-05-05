package main

import "fmt"

func do(s string) string {
	char := make([]int32, 0, len(s)/2+1)
	num := make([]int32, 0, len(s)/2+1)
	res := make([]int32, 0, len(s))
	for _, val := range s {
		if val >= 48 && val <= 57 {
			num = append(num, val)
		} else {
			char = append(char, val)
		}
	}

	charNum := len(char)
	numNum := len(num)

	if charNum-numNum > 1 || charNum-numNum < (-1) {
		return ""
	}

	if charNum > numNum {
		for len(num)+len(char) > 0 {
			if len(char) > len(num) {
				res = append(res, char[0])
				char = char[1:]
			} else {
				res = append(res, num[0])
				num = num[1:]
			}
		}
	} else {
		for len(num)+len(char) > 0 {
			if len(num) > len(char) {
				res = append(res, num[0])
				num = num[1:]
			} else {
				res = append(res, char[0])
				char = char[1:]
			}
		}
	}
	return string(res)
}

func main() {
	//str := "09abc"
	//for _, val := range str {
	//	fmt.Println(val)
	//}
	s := "aaaaa1111"
	fmt.Println(do(s), "/")
}
