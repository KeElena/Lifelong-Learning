func maxLengthBetweenEqualCharacters(s string) int {
    
    max:=-1                                         //设置默认为-1
    store:=make([][]int,26)                         //构建存储索引的二维数组
    for i,char:=range s{                            //遍历字符串并追加索引
        store[char-97]=append(store[char-97],i)
    }

    for _,s:=range store{                           //遍历存储二维数组
        if len(s)<=1{                               //跳过长度小于2的行
            continue
        }

        if s[len(s)-1]-s[0]-1>max{                  //计算最后一个索引和第1个索引的差值-1
            max=s[len(s)-1]-s[0]-1                  //大于max则赋值
        }
    }

    return max                                      //返回结果
}