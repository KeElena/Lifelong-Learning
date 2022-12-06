package main

//用于求1000量级的组合数
//用于绕过阶乘求组合数的高效方法（主要运算是加分不是乘法）
//先构建杨辉三角，根据杨辉三角的特性进行递归运算，输出杨辉三角的第n列（从0开始）
func cal(n,m int) int {
    triAngle:=buildAngle(n+1,[]int{})
    return triAngle[m]
}
//需要的参数：num、last
//num：杨辉三角的行数，求组合数要求n+1行
//last：杨辉三角上一行的数组，开始默认为[]int{}
func buildAngle(num int,last []int)(res []int){
    
    if len(last)==num{                                      //last数组的长度对应当前位于杨辉三角的第len行，达到后返回last
        return last
    }
    res=make([]int,0,len(last)+1)                           //初始化数组
    for i:=0;i<len(last)+1;i++{                             //for构建杨辉三角的下一行
        if i==0 ||i==len(last){                             //第一个元素和最后一个元素默认为1
            res=append(res,1)
        }else{
            res=append(res,((last[i]+last[i-1])%1000000007))  //其他元素为上一行的第i个元素和第i-1个元素的和，由于阶数爆炸需要取模
        }
    }
    res=buildAngle(num,res)                                 //新构建行的传入递归函数，返回杨辉三角的第num行
    return
}

func main(){
    fmt.Println(cal(500,45))
}