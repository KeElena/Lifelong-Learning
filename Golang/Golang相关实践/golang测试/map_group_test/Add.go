package test_group_map

func Add(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}
