//使用散列表记录每个字符出现的位置
func uniqueLetterString(s string) int {
	indexMap := make(map[rune][]int)							//对应散列表
	var res int
	for i, char := range s {									//遍历字符串并记录对应字母出现的位置
		indexMap[char] = append(indexMap[char], i)
	}

	for _, arr := range indexMap {								//遍历散列表
		arr = append([]int{-1}, append(arr, len(s))...)			//给每个切片追加前缀和后缀
		for i := 1; i < len(arr)-1; i++ {						//遍历数组，i从1开始到len(arr)-1结束，防止追加的前后缀对结果的干扰
			res += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])	//计算并累加结果
		}
	}
	return res
}