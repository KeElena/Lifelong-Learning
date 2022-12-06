package test_group

func Add(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
