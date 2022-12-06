func rotatedDigits(n int) int {
    res:=0
    var str string
    var sum int
    for i:=2;i<=n;i++{
        str=strconv.Itoa(i)
        sum=0
        for _,char:=range str{
            if char-48!=2 &&char-48!=5 &&char-48!=6&&char-48!=9{
                break
            }
            sum++
        }

        if sum==len(str){
            res++
        }
    }
    return res
}