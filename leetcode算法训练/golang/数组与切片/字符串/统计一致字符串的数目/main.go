func countConsistentStrings(allowed string, words []string) int {
    //初始化散列表和相关参数
    index:=make(map[rune]bool)
    var sum int
    //初始化散列表的值
    for _,char:=range allowed{
        index[char]=true
    }
    //遍历所有字符串
    for _,str:=range words{
        //遍历某一字符串内的所有字符
        for _,char:=range str{
            //如果散列表中字符不存在则计数+1，并退出扫描该字符串的循环
            if index[char]==false{
                sum++
                break
            }
        }
    }
    return len(words)-sum
}