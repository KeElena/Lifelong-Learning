func rev(x int)(res int){
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}