func explorationSupply(station []int, pos []int) []int {
    
    var distance int
    var idx int
    res:=make([]int,0,len(pos))
    for _,p:=range pos{
        distance=100000000              //设距离初始值为最大值或最大值范围之外
        idx=-1                          //设置位置索引的初始值为-1
        for i,val:=range station{       //遍历补给站
            if cal(p,val)<distance{     //勘探队与补给站的距离小于distance时
                distance=cal(p,val)     //distance赋值
                idx=i                   //索引赋值
            }
        }
        res=append(res,idx)             //遍历了后获取了最小距离的索引，追加该索引
    }
    
    return res                          //返回结果
    
}

func cal(a,b int)int{                   //计算相减的绝对值
    if a-b>0{
        return a-b
    }
    return b-a
}