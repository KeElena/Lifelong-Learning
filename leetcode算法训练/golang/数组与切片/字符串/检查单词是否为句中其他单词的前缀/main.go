package main

import (
	"fmt"
	"strings"
)

func do(sentence string, searchWord string) int {

	charList := strings.Split(sentence, " ")				//以空格作为分割符将字符串分割
	for index, word := range charList {						//遍历每一个单词
		if len(word) >= len(searchWord) && word[:len(searchWord)] == searchWord {
			return index + 1								//要求单词的长度要大于等于前缀的长度
		}													//截取单词与前缀等长前几个字符进行比较
	}
	return -1
}

func main() {
	str := "this problem is an easy problem"
	search := "pro"
	fmt.Println(do(str, search))
}
