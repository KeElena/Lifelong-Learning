func scoreOfParentheses(s string) int {
    res:=make([]int,len(s))
    var p int
    for _,char:=range s{
        if char=='('{
            p++
            res[p]=0
            continue
        }
        if char==')'{
            p--
            res[p]+=getMax(2*res[p+1],1)
            continue
        }
    }
    return res[0]
}
func getMax(i,j int)int{
    if i>j{
        return i
    }
    return j
}