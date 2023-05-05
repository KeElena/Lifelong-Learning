func canFormArray(arr []int, pieces [][]int) bool {
    table:=make(map[int][]int,len(pieces))          //构建map表
    for _,s:=range pieces{                          //以每个数组的第一个元素为key存储数组
        table[s[0]]=s
    }

    for len(arr)!=0{                                //长度非0时循环
        s:=table[arr[0]]                            //根据arr[0]的值在map表里取数组
        if s==nil{                                  //s为nil时返回false
            return false
        }
        for i,val:=range s{                         //遍历s
            if arr[i]!=val{                         //根据索引对比arr的前几个元素
                return false                        //如果不同则返回false
            }
        }
        arr=arr[len(s):]                            //裁剪arr直到len(arr)==0时跳出循环
    }
    return true                                     //跳出循环默认返回true
}