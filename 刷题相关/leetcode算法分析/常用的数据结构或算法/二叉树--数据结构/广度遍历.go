语言： golang

添加备注


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
    layor=1                         //默认第一层
    Queue:=make([]*TreeNode,0,5)
    Queue=append(Queue,root)

    for len(Queue)!=0{
        
        size=len(Queue)
        for _,node:=range Queue{
            if node.Left!=nil{
                Queue=append(Queue,node.Left)
            }
            if node.Right!=nil{
                Queue=append(Queue,node.Right)
            }
        }
        
        Queue=Queue[size:]        
        layor++
    }
    return root
}