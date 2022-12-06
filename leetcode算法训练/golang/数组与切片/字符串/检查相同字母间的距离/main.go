func checkDistances(s string, distance []int) bool {
    Map:=make(map[string][]int,len(s)/2)                //构造一个散列表
    ok:=true                                           //定义和默认值为true
    
    for i:=range s{                                     //遍历字符串
        Map[s[i:i+1]]=append(Map[s[i:i+1]],i)           //字符串切割成单个字符作为key
    }
    
    for key,arr:=range Map{
        Byte:=[]byte(key)                               //将string转为字节（即ASCII码）
        if arr[1]-arr[0]-1 != distance[Byte[0]-97]{     //切片的索引可以是uint8
            ok=ok && false                            //逻辑的与运算
        }
    }
    return ok
}