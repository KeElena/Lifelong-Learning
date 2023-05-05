type OrderedStream struct {
    cap int
    ptr int
    stream map[int]string
}

func Constructor(n int) OrderedStream {
    return OrderedStream{cap:n,ptr:1,stream:make(map[int]string,n)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.stream[idKey]=value
    list:=[]string{}
    if idKey == this.ptr{
        for i:=this.ptr;i<this.cap+1;i++{
            if this.stream[i]=="" || this.ptr>this.cap{
                return list
            }
            list=append(list,this.stream[i])
            this.ptr++
        }
    } 
    return list
}