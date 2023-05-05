func minMaxGame(nums []int) int {
    if len(nums)==1{
        return nums[0];
    }
    arr:=make([]int,len(nums)/2)

    for i:=range arr{
        if i%2==0{
            arr[i]=getMin(nums[2*i],nums[2*i+1])
        }else{
            arr[i]=getMax(nums[2*i],nums[2*i+1])
        }
    }
    return minMaxGame(arr)
}
func getMin(a,b int)int{
    if a<b{
        return a
    }
    return b
}
func getMax(a,b int)int{
    if a>b{
        return a
    }
    return b
}