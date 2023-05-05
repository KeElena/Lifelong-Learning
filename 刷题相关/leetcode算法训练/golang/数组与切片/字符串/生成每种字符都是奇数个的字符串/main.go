package main

import (
	"fmt"
	"strings"
)

func do(n int) string {

	if n%2 == 0 {								//n为偶数时
		return "b" + strings.Repeat("a", n-1)	//加1个a，其他全b
	}
	return strings.Repeat("a", n)				//n为奇数时全a
}
func main() {
	fmt.Println(do(8))
}
