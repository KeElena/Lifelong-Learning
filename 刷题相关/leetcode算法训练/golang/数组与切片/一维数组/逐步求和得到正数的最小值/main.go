package main

import "fmt"

func do(nums []int) int {
	i, sum, bol := 1, 0, true
	for {
		bol = true
		sum += i
		for _, val := range nums {
			sum += val
			if sum < 1 {
				bol = bol && false
				break
			}
		}

		if bol == true {
			return i
		}
		sum -= sum
		i++
	}
	return 0
}
func main() {
	arr := []int{1, -2, -3}
	fmt.Println(do(arr))
}
