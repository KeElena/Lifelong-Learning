//存储相同字母之间的距离，通过距离直接计算结果
//需要切片记录上一个字母出现的位置
func uniqueLetterString(s string) int {
	indexList := make([][]int, 26)					//用于存储距离
	lastIndex := make([]int, 26)					//用于存储字母上次出现的位置
	var res int
	for i, char := range s {						//遍历字符串
		if len(indexList[char-65]) == 0 {			//存储距离的二维切片中，如果对应字母切片的长度为0时，说明是第一次遇到的字母，否则就是多次遇到的字母
			indexList[char-65] = append(indexList[char-65], i+1)	//默认最左边界为-1，即(i-(-1))
		} else {
			indexList[char-65] = append(indexList[char-65], i-lastIndex[char-65])	//如果是多次遇到的字母则(当前索引-字母上次出现的索引)
		}
		lastIndex[char-65] = i						//存储该字母的索引，记录该字母已出现以及其位置
	}

	for index, arr := range indexList {				//遍历存储距离的二维数组
		if len(arr) == 0 {							//字母从没出现过则跳过
			continue
		}
		arr = append(arr, len(s)-lastIndex[index])	//追加边界与字母最后出现的索引的距离

		for i := 0; i < len(arr)-1; i++ {			//遍历所有距离，当前索引的距离与下一个距离相乘并累加到结果里
			res += arr[i] * arr[i+1]				//要求i要小于len(arr)-1，否则会超出循环范围
		}
	}
	return res
}