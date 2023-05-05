package main

import (
	"fmt"
	"strings"
)

func addSuffix(text string, suffix string) func() string {
	f := func() string {
		if strings.HasSuffix(text, suffix) {	//判断文本是否含后缀
			return text			//有则直接返回
		} else {
			return text + suffix		//无则添加后缀返回
		}
	}
	return f					//返回函数
}

func main() {
	text := "work"
	fmt.Println(addSuffix(text, ".jpg")())		//返回函数并执行
}
