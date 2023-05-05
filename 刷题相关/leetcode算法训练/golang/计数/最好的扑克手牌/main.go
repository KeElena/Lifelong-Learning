func bestHand(ranks []int, suits []byte) string {
    var rMax int
    var sMax int
    rankMap:=make(map[int]int,len(ranks))
    suitMap:=make(map[byte]int,len(suits))

    for _,v:=range ranks{
        rankMap[v]++
        rMax=getMax(rMax,rankMap[v])
    }
    for _,v:=range suits{
        suitMap[v]++
        sMax=getMax(rMax,suitMap[v])
    }
    if sMax==5{
        return "Flush"
    }else if rMax>=3{
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