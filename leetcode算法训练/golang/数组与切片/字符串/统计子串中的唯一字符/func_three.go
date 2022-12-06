//由于散列表存取数据比切片耗时，使用二维切片代替散列表可以加快运算速度
func uniqueLetterString(s string) int {
	indexList := make([][]int, 26)							//定义有26行的二维切片，对应26个字母
	var res int
	for i, char := range s {
		indexList[char-65] = append(indexList[char-65], i)	//遍历字符串，(uncode编码-65)可以映射到二维切片对应的行并保留字母顺序
	}

	for _, arr := range indexList {
		if len(arr) == 0 {									//数组长度为0时直接跳过
			continue
		}
		arr = append([]int{-1}, append(arr, len(s))...)		//追加前缀和后缀，下面和func_two处理的思路一样
		for i := 1; i < len(arr)-1; i++ {
			res += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return res
}