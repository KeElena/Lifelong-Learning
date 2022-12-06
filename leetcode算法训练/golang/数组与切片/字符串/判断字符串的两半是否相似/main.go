func halvesAreAlike(s string) bool {
    //构建散列表
    m:=make(map[byte]bool,10)
    var a,b int
    m['a']=true
    m['e']=true
    m['i']=true
    m['o']=true
    m['u']=true
    m['A']=true
    m['E']=true
    m['I']=true
    m['O']=true
    m['U']=true
    //遍历字符串计数
    for i:=0;i<len(s);i++{
        if m[s[i]]{
            if i<len(s)/2{
                a++
            }else{
                b++
            }
        }
    }
    //返回结果
    return a==b
}