type OrderedStream struct {
    ptr int
    stream []string
}

func Constructor(n int) OrderedStream {
    return OrderedStream{ptr:0,stream:make([]string,n,n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.stream[idKey-1]=value
    list:=[]string{}
    if idKey-1 == this.ptr{
        for i:=this.ptr;i<len(this.stream);i++{
            if this.stream[i]=="" || this.ptr>len(this.stream){
                return list
            }
            list=append(list,this.stream[i])
            this.ptr++
        } 
    }
    return list
}
