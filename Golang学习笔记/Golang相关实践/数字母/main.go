package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "how do you do you"
	slice := strings.Split(str, " ")
	fmt.Println(slice)
	m1 := make(map[string]int, 10)

	for i := 0; i < len(slice); i++ {
		m1[slice[i]] = m1[slice[i]] + 1
	}
	for key, val := range m1 {
		fmt.Printf("%s的个数%d\n", key, val)
	}
}
