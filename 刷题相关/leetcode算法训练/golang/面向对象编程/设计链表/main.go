type MyLinkedList struct {
    Val int
    next *MyLinkedList
}

func Constructor() MyLinkedList {
    return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
    node:=this
    for i:=0;i<index+1;i++{
        if node.next==nil{
            return -1
        }
        node=node.next
    }

    return node.Val
}

func (this *MyLinkedList) AddAtHead(val int)  {
    this.next=&MyLinkedList{Val:val,next:this.next}
}

func (this *MyLinkedList) AddAtTail(val int)  {
    node:=this
    for node.next!=nil{
        node=node.next
    }
    node.next=&MyLinkedList{Val:val}
}

func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    node:=this
    for i:=0;i<index;i++{
        if node.next==nil{
            return
        }
        node=node.next
    }
    node.next=&MyLinkedList{Val:val,next:node.next}
}

func (this *MyLinkedList) DeleteAtIndex(index int){
    node:=this
    for i:=0;i<index;i++{
        if node.next==nil{
            return
        }
        node=node.next
    }
    tempNode:=node.next
    if tempNode==nil{
        return
    }
    node.next=tempNode.next
    tempNode.next=nil
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */