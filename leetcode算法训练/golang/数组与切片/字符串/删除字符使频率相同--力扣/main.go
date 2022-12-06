func equalFrequency(word string) bool {
    var num int
    var ok bool
    var count map[rune]int
    for i:=0;i<len(word);i++{
        count=make(map[rune]int)        
        for j,val:=range word{
            if j==i{
                continue
            }
            count[val]++
        }
        num=-1
        ok=true
        for _,val:=range count{
            if num==-1{
                num=val
                continue
            }
            if num!=val{
                ok=false
                break
            }
        }
        if ok==true{
            return true
        }
    }
    return false
}