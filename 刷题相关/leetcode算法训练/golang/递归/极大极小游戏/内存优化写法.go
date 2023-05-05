func minMaxGame(nums []int) int {
    if len(nums)==1{
        return nums[0];
    }


    for i:=0;i<len(nums)/2;i++{
        if i%2==0{
            nums[i]=getMin(nums[2*i],nums[2*i+1])
        }else{
            nums[i]=getMax(nums[2*i],nums[2*i+1])
        }
    }
    return minMaxGame(nums[:len(nums)/2])
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