func shuffle(nums []int, n int) []int {
    x:=nums[:len(nums)/2]               //截取x数组
    y:=nums[len(nums)/2:]               //截取y数组
    res:=make([]int,0,len(nums))        //初始化res数组

    for i:=0;i<len(nums);i++{
        if i%2 ==0{                     //偶数位元素为x的元素
            res=append(res,x[0])        //追加x的第一个元素
            x=x[1:]                     //剔除第一个元素
        }
        if i%2 ==1{                     //奇数位元素为y的元素
            res=append(res,y[0])        //追加y的第一个元素
            y=y[1:]                     //剔除第一个元素
        }
    }
    return res
}