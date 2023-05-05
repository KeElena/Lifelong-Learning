func checkOnesSegment(s string) bool {
    var num int                         //统计连续1的字段数
    var count int                       //统计字段里1的个数
    s+="0"                              //添加后缀0
    for _,char:=range s{                //遍历s
        if char==49{                    //char为1时，1的个数增加
            count++
        }else{
            if count>0{                 //char为0时，如果count>0则字段数+1
                num++
            }
            count=0                     //重置计数
        }
        if num>1{                       //num>1时不满足条件，返回false
            return false
        }
    }
    return true
}