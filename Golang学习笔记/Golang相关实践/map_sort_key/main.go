package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	m1 := make(map[string]int, 20)
	for i := 0; i < len(m1); i++ {
		key := fmt.Sprintf("stu%03d", rand.Intn(200)) //取随机数时可能存在相同数
		value := rand.Intn(100)
		m1[key] = value
	}

	keys := make([]string, 0)
	for key := range m1 {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("Sno=%s val=%d\n", key, m1[key])
	}
}
