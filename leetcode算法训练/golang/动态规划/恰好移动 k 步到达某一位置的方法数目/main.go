func numberOfWays(startPos int, endPos int, k int) int {
   var move,res int
   move=endPos-startPos                 //到达目的需要移动的位移
    triAngle:=buildAngle(k+1,[]int{})   //构建杨辉三角，获取第k+1行的数组

    for i:=0;i<=k;i++{                  //遍历右位移i的所有可能取值
        if i==k-i+move{                 //要求右位移=左位移+需要移动的位移
            res+=triAngle[i]%1000000007 //满足条件直接取第i个元素进行求和取模
        }
    }
    return res
}
//杨辉三角求组合数
func buildAngle(num int,last []int)(res []int){
    
    if len(last)==num{
        return last
    }
    res=make([]int,0,len(last)+1)
    for i:=0;i<len(last)+1;i++{
        if i==0 ||i==len(last){
            res=append(res,1)
        }else{
            res=append(res,((last[i]+last[i-1])%1000000007))
        }
    }
    res=buildAngle(num,res)
    return
}