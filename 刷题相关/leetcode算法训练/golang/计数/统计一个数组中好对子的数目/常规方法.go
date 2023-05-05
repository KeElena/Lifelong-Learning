func countNicePairs(nums []int) int {
    var count int
    //记忆化存储
    tempArr:=make([]int,0,len(nums))
    //获取所有数的反转值
    for _,v:=range nums{
        tempArr=append(tempArr,rev(strconv.Itoa(v)))
    }
    //双循环遍历
    for i,x:=range nums{
        for j:=i+1;j<len(nums);j++{
            if x+tempArr[j]==tempArr[i]+nums[j]{
                count++;
            }
        }
    }
    return count;
}
//不推荐的反转方式
func rev(num string)(result int){
    var temp string
    for i:=len(num)-1;i>-1;i--{
        temp=temp+num[i:i+1]
    }
    result,_=strconv.Atoi(temp)
    return
}