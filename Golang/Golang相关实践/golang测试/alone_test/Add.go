package test_one

func Add(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
