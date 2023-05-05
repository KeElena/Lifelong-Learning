func bestHand(ranks []int, suits []byte) string {
    var rMax int
    isFiush:=true
    rankMap:=make(map[int]int,len(ranks))

    for i:=1;i<len(suits);i++{
        if suits[i]!=suits[i-1]{
            isFiush=false
            break
        }
    }
    if isFiush{
        return "Flush"
    }

    for _,v:=range ranks{
        rankMap[v]++
        rMax=getMax(rMax,rankMap[v])
    }

    if rMax>=3{
        return "Three of a Kind"
    }else if rMax==2{
        return "Pair"
    }else if rMax==1{
        return "High Card"
    }
    return ""
}
func getMax(a,b int)int{
    if a>b{
        return a
    }
    return b
}