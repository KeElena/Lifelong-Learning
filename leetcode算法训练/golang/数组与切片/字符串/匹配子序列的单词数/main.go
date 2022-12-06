func numMatchingSubseq(s string, words []string) int {
    //初始化字母表，切片存储words的idx索引（通过索引定位字符串）
    m:=make([][]int,26)
    var num int
    //根据字符串的首字母进行分类
    for i,str:=range words{
        m[str[0]-97]=append(m[str[0]-97],i)
    }
    //遍历给的的字符串
    for _,char:=range s{
        //如果以当前字符为首字母的字符串不在words中则跳过
        if len(m[char-97])==0{
            continue
        }
        //初始化缓存切片
        var temp []int
        //遍历以该字符为首字母的字符串（获取idx，通过idx在words中访问字符串）      
        for _,idx:=range m[char-97]{
            //去除首字母
            words[idx]=words[idx][1:]
            //如果字符串长度为0，则计数++，跳过循环
            if len(words[idx])==0{
                num++
                continue
            }
            //如果去除首字母后还是以当前字符为首字母
            if words[idx][0]-97==byte(char-97){
                //将索引加到缓存切片
                temp=append(temp,idx)
            }else{
                //不是则将索引加到字母表对应字母的索引切片
                m[words[idx][0]-97]=append(m[words[idx][0]-97],idx)
            }
        }
        //用缓存切片更新当前字符的索引切片
        m[char-97]=temp
    }
    return num
}