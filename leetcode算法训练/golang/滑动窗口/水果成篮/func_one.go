func totalFruit(fruits []int) int {
    var max int
    var left int
    count:=make(map[int]int,3)
    for rignt,val:=range fruits{
        count[val]++
        for len(count)>2{
            count[fruits[left]]--
            if count[fruits[left]]==0{
                delete(count,fruits[left])
            }
            left++
        }
        max=getMax(max,rignt-left+1)
    }
    return max
}
func getMax(a,b int)int{
    if a>b{
        return a
    }
    return b
}