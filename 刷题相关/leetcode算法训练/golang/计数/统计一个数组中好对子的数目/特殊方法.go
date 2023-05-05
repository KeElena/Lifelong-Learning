func countNicePairs(nums []int) int {
    var count int
    //初始化哈希
    m:=make(map[int]int,5)
    var mod int=1e9+7
    //对式子进行变换，同变量放在一边
    for _,v:=range nums{
        m[v-rev(v)]++
    }
    //使用特殊组合数进行累加，C^2_n
    for _,v:=range m{
        count+=v*(v-1)/2
    }
    return count %mod;
}
//反转
func rev(x int)(res int){
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}