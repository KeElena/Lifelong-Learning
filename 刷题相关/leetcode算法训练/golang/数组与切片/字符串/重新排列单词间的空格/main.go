package main

import (
	"fmt"
	"strings"
)

func do(text string) string {

	num := strings.Count(text, " ")						//使用strings.Count()方法对空格进行计数
	words := strings.Fields(text)						//使用strings.Fields()方法将单词取出存在切片中
	if len(words) == 1 {								//如果只有一个单词则字节将空格查到单词的后面然后返回
		return words[0] + strings.Repeat(" ", num)
	}
	text = ""											//text清空
	var val int
	val = num / (len(words) - 1)						//计算每个单词的间隔里空格的数量，要求words长度要大于1，整形除法会舍去尾数
	for i, str := range words {							//遍历所有单词
		text += str										//单词追加到text
		if i == len(words)-1 {							//如果遍历到最后一个单词则追加剩余的空格
			text += strings.Repeat(" ", num)
		} else {
			text += strings.Repeat(" ", val)			//追加val个空格
			num -= val									//减去已使用的空格数量
		}
	}
	return text
}

func main() {
	str := " a"
	fmt.Print(1, do(str), 1)
}
