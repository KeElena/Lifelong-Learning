func sumPrefixScores(words []string) []int {
    res:=make([]int,len(words))                                 //返回结果
    var sum int                                                 //累加缓存
    var set []string                                            //存储前缀
    for i,word:=range words{                                    //循环每个字符串
        for i:=1;i<=len(word);i++{
            set=append(set,word[:i])                            //提取前缀
        }

        for _,suf:=range set{                                   //遍历提取出的前缀
            for _,word:=range words{                            //遍历每个字符串
                if len(suf)<=len(word) && word[:len(suf)]==suf{ //前缀匹配
                    sum++                                       //计数+1
                }
            }
        }

        res[i]=sum                                              //遍历完第一个字符串输出结果
        set=[]string{}                                          //重置前缀
        sum=0                                                   //重置计数缓存
    }
    return res
}