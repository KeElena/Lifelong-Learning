func areAlmostEqual(s1 string, s2 string) bool {
    temp:=make([]int,0,2)
    for i:=0;i<len(s1);i++{
        if s1[i]!=s2[i]{
            if len(temp)<2{
                temp=append(temp,i)
            }else{
                return false
            }
        }
    }
    if len(temp)==0 || (len(temp)==2&&s1[temp[0]]==s2[temp[1]]&&s1[temp[1]]==s2[temp[0]]) {
        return true
    }
    return false
}