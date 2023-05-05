func CheckPermutation(s1 string, s2 string) bool {
    PC:=make(map[rune]int)      //计数器
    for _,char:=range s1{       //计数+1
        PC[char]++
    }
    for _,char:=range s2{       //计数-1
        PC[char]--
    }
    for _,val:=range PC{        //如果val非0说明不能重排变成另一个字符串
        if val!=0{
            return false
        }
    }
    return true
}