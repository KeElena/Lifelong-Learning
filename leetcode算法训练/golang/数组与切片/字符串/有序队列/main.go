package main

import (
	"fmt"
	"sort"
)

func do(s string, k int) (minStr string) {
	if k == 1 {
		minStr = s

		for i := 0; i < len(s); i++ {
			s = s[1:] + s[:1]
			if s < minStr {
				minStr = s
			}
		}
		return
	}

	cp := []byte(s)
	sort.Slice(cp, func(i, j int) bool {
		return cp[i] < cp[j]
	})

	return string(cp)
}

func main() {
	str := "xitavoyjqiupzadbdyymyvuteolyeerecnuptghlzsynozeuuvteryojyokpufanyrqqmtgxhyycltlnusyeyyqygwupcaagtkuq"
	k := 1
	fmt.Println(do(str, k))
}
