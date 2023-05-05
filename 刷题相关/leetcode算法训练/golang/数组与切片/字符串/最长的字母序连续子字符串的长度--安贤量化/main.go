func longestContinuousSubstring(s string) int {
    var max int
    var sum int
    sum=1
    max=1
    for i:=1;i<len(s);i++{
        if i==0{
            continue
        }
        if s[i]-s[i-1]==1{
            sum+=1
            max=getMax(max,sum)           
        }else{
            sum=1     
        }
    }
    return max
}
func getMax(i,j int)int{
    if i>j{
        return i
    }
    return j
}