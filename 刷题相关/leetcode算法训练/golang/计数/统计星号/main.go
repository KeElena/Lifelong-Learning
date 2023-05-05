func countAsterisks(s string) int {
    var count int
    //定义判断标志变量
    var hasLeft bool
    for _,v:=range s{
        //如果v为|
        if v==124{
            //hasLeft为false时标志变量值为true
            if !hasLeft{
                hasLeft=true
            }else{
            //如果有左|则标志变量设为false
                hasLeft=false
            }
            continue
        }
        //在非累计区域则跳过
        if hasLeft{
            continue
        }
        //值为*的ASCII码时累加
        if v==42{
            count++
        }
    }
    return count
}