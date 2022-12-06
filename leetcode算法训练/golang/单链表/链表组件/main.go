/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	m:=make(map[int]bool,len(nums))
    var num int
	for _, v := range nums {
		m[v] = true
	}

	for {
        if m[head.Val]==true && head.Next==nil{
            num++
            break
        }
        if head.Next==nil{
            break
        }
        if m[head.Val]==true && m[head.Next.Val]==false{
            num++
        }
        head=head.Next
	}
	return num
}