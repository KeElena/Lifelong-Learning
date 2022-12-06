type LUPrefix struct {
    idx []byte
    i int
    count int
}

func Constructor(n int) LUPrefix {
    return LUPrefix{idx:make([]byte,n)}
}

func (this *LUPrefix) Upload(video int)  {
    this.idx[video-1]=1
}

func (this *LUPrefix) Longest() int {
    
    for i:=this.count;i<len(this.idx);i++{
        if this.idx[i]==1{
            this.count++
        }else{
            break
        }
    }
    return this.count
}
