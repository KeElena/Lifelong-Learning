func validateStackSequences(pushed []int, popped []int) bool {
    arr:=make([]int,0,len(pushed))
    index:=0
    for i:=0;i<len(pushed);i++{
        if index!=len(pushed){                          //1个1个入栈
            arr=append(arr,pushed[index])
            index++
        }
        
        for len(arr)>0 && popped[0]==arr[len(arr)-1]{   //for循环在满足条件下可以一次出栈多个元素
            arr=arr[:len(arr)-1]
            popped=popped[1:]
        }
    }

    if len(popped)==0{
        return true
    }
    return false
}