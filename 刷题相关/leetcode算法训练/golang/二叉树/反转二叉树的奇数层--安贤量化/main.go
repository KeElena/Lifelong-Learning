/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    
    var size,layor int
    var arr []int
    layor=1
    Queue:=make([]*TreeNode,0,5)
    Queue=append(Queue,root)
    for len(Queue)!=0{
        
        size=len(Queue)
        for _,node:=range Queue{
            if node.Left!=nil{
                Queue=append(Queue,node.Left)
                arr=append(arr,node.Left.Val)
            }
            if node.Right!=nil{
                Queue=append(Queue,node.Right)
                arr=append(arr,node.Right.Val)
            }
        }
        
        Queue=Queue[size:]        
        if layor%2==1{
            for i,node:=range Queue{
                node.Val=arr[len(arr)-1-i]
            }
        }
        arr=[]int{}
        layor++
    }
    return root
}