func reContruct(head *ListNode) *ListNode {
    if head==nil{
        return nil
    }

    head.Next=reContruct(head.Next)
    
    return head
}