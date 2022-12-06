func scoreOfParentheses(s string) int {
    res:=make([]int,len(s))
    var p int
    for i,char:=range s{
        if char=='('{
            p++
            res[p]=0
            continue
        }
        if char==')' && s[i-1]=='('{
            p--
            res[p]+=1
            continue
        }
        if char==')' && s[i-1]==')'{
            p--
            res[p]+=2*res[p+1]
        }
    }
    return res[0]
}