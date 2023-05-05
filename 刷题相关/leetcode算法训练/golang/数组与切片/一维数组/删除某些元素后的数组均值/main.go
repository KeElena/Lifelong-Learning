func trimMean(arr []int) float64 {
    sort.Ints(arr)                          //给arr排序
    var res float64
    n:=5*len(arr)/100                       //一边需要删除元素的个数
    
    for i:=n;i<len(arr)-n;i++{              //从第n+1个元素开始遍历，到len(arr)-n-1位置的元素停止
        res+=float64(arr[i])                //累加
        if i==len(arr)-n-1{
            res=res/float64(len(arr)-2*n)   //累加完后求平均
        }
    }
    
    return res
}