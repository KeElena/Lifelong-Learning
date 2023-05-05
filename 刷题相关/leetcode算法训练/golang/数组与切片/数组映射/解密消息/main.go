func decodeMessage(key string, message string) string {
    //定义出现判断数组
    appear:=make([]bool,26)
    //定义映射数组
    store:=make([]byte,26)
    //字母表索引，用byte方便计算
    var charIdx byte
    //key索引，不能用byte否则死循环
    var keyIdx int
    for charIdx<26{
        //判断不是空格和是没出现的字符
        if key[keyIdx]!=' '&&!appear[key[keyIdx]-97]{
            //设置映射字母
            store[key[keyIdx]-97]=97+charIdx
            //记录已出现
            appear[key[keyIdx]-97]=true
            //已映射字母+1
            charIdx++
        }
        //key索引+1
        keyIdx++
    }
    //使用字节数组修改字符串
    temp:=[]byte(message)
    for i,char:=range temp{
        //跳过空格
        if char==32{
            continue
        }
        //根据字母获取映射字母
        temp[i]=store[char-97]
    }
    //返回密钥
    return string(temp)
}