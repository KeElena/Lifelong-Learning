//双端队列结构
type MyCircularDeque struct {
    Queue []int           //队列主体
    num int               //队列容量
}
//初始化队列
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{Queue:make([]int,0,k),num:k}         //初始化并返回队列
}
//前插
func (this *MyCircularDeque) InsertFront(value int) bool {
    if len(this.Queue)==this.num{
        return false
    }
    this.Queue=append([]int{value},this.Queue...)               //通过append实现前插
    return true
}
//尾插
func (this *MyCircularDeque) InsertLast(value int) bool {
    if len(this.Queue)==this.num{
        return false
    }
    this.Queue=append(this.Queue,value)                         //通过append实现尾插
    return true
}
//前删
func (this *MyCircularDeque) DeleteFront() bool {
    if len(this.Queue)==0{
        return false
    }
    this.Queue=this.Queue[1:]
    return true
}
//尾删
func (this *MyCircularDeque) DeleteLast() bool {
    if len(this.Queue)==0{
        return false
    }
    this.Queue=this.Queue[:len(this.Queue)-1]
    return true
}
//取队头
func (this *MyCircularDeque) GetFront() int {
    if len(this.Queue)==0{
        return -1
    }
    return this.Queue[0]
}
//取队尾
func (this *MyCircularDeque) GetRear() int {
    if len(this.Queue)==0{
        return -1
    }
    return this.Queue[len(this.Queue)-1]
}
//判断空
func (this *MyCircularDeque) IsEmpty() bool {
    if len(this.Queue)==0{
        return true
    }
    return false
}
//判断满
func (this *MyCircularDeque) IsFull() bool {
    if len(this.Queue)==this.num{
        return true
    }
    return false
}