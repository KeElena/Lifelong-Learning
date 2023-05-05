var diffs = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][2]int, m)
	for i := range dp {
		dp[i] = [2]int{-1, -1}
	}
	var f func(int, int, bool) int
	f = func(i, isDiff int, isLimit bool) (res int) {
		if i == m {
			return isDiff // 只有包含 2/5/6/9 才算一个好数
		}
		if !isLimit {
			dv := &dp[i][isDiff]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for d := 0; d <= up; d++ { // 枚举要填入的数字 d
			if diffs[d] != -1 { // d 不是 3/4/7
				res += f(i+1, isDiff|diffs[d], isLimit && d == up)
			}
		}
		return
	}
	return f(0, 0, true)
}