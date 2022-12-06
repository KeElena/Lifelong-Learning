//暴力解法，容易超时
func uniqueLetterString(s string) int {
	res := len(s)
	//使用滑块截取字符串
	for size := 2; size <= len(s); size++ {	//size为滑块大小
		j := 0								//滑块左游标默认为0
		for i := size; i <= len(s); i++ {	//右游标默认等于size
			res += count(s[j:i])			//截取字符串使用count函数统计唯一字符数量，返回唯一字符个数并累加
			j++								//左右向右滑动，i因为循环默认自增，需要补充j的自增
		}
	}

	return res								//返回结果
}

func count(str string) (res int) {
	count := make([]int, 26)			//默认有26个字母

	for _, ascii := range []byte(str) {	//字符串转为字节切片，遍历并统计字符出现的次数
		count[ascii-65] += 1
	}
	for _, num := range count {			//统计出现次数为1的字符数
		if num == 1 {
			res += 1
		}
	}
	return								//返回结果
}