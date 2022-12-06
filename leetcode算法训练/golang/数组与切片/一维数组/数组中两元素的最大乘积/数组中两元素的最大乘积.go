func maxProduct(nums []int) int {
    index,max:=0,0                  //定义参数
    res:=make([]int,0,2)            //构造切片
    for i:=0;i<2;i++{               //获取两个最大值
        index,max=0,0               //重置参数
        for i,val:=range nums{      //遍历数组
            if val>max{             //获取最大值
                max=val
                index=i
            }
        }
        res=append(res,max-1)       //追加元素
        if i!=1{
            nums=append(nums[:index],nums[index+1:]...)     //剔除最大值对应的元素            
        }
    }
    return res[0]*res[1]
}