func minAddToMakeValid(s string) int {
    var count int
    var res int
    for _,char:=range s{
        if char=='('{
            count++
            continue
        }
        if char==')'{
            count--
            if count<0{
                res++
                count=0
            }
        }
    }
    return res+count
}