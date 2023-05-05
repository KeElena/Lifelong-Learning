func decrypt(code []int, k int) []int {
    res:=make([]int,len(code))
    var sum int

    if k==0{
        return res
    }

    if k>0{
        for i:=0;i<len(code);i++{
            sum=0
            for j:=1;j<=k;j++{
                sum+=code[(i+j)%len(code)]
            }
            res[i]=sum
        }
    }else{
        for i:=0;i<len(code);i++{
            sum=0
            for j:=1;j<=-k;j++{
                sum+=code[(i+len(code)-j)%len(code)]
            }
            res[i]=sum
        }
    }
    return res
}