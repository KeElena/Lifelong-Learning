func repeatedCharacter(s string) byte {
    count:=make(map[byte]int,len(s))
    for i:=0;i<len(s);i++{
        count[s[i]]++
        if count[s[i]]==2{
            return s[i]
        }
    }
    return 0
}