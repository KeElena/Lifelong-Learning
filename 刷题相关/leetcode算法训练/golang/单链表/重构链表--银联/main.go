/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reContruct(head *ListNode) *ListNode {
    if head==nil{                       //结点为nil时返回nil
        return nil
    }

    head.Next=reContruct(head.Next)     //head.Next获取返回的地址
    
    if head.Val%2==0{                   //当前节点被2整除时返回head.Next结点
        return head.Next
    }
    return head                         //默认返回head结点
}