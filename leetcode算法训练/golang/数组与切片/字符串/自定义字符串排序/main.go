func customSortString(order string, s string) string {
    //定义字母顺序表
    table:=make(map[rune]int,len(order))
    //定义字符串字母顺序数组
    arr1:=make([]int,0,len(s))
    //定义暂存字节切片
    arr2:=make([]byte,0,len(s))
    //定义结果字节切片
    result:=make([]byte,0,len(s))
    //位置顺序以1开始进行记录
    for i,char:=range order{
        table[char]=i+1
    }
    //遍历字符串
    for _,char:=range s{
        //字母顺序表输出非0时说明字母存在顺序
        if table[char]!=0{
            //暂存到字母顺序数组
            arr1=append(arr1,table[char])
        }else{
            //没有顺序限制的字母直接以字节的形式追加到暂存字节切片，rune类型需要转为byte类型
            arr2=append(arr2,byte(char))
        }
    }
    //对顺序数组进行排序
    sort.Ints(arr1)
    //根据顺序编号从order字符串里取字母到结果切片
    for _,v:=range arr1{
        result=append(result,order[v-1])
    }
    //将暂存字节切片的字母追加到结果切片
    result=append(result,arr2...)
    //返回字符串
    return string(result)
}