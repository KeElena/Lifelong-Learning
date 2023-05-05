package main

import "fmt"

//树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//广度遍历结果存储结构
type CBTInserter struct {
	Root  *TreeNode
	Queue []*TreeNode
}

//进行广度遍历并存储地址
func Constructor(root *TreeNode) CBTInserter {

	var CBTI = CBTInserter{Root: root}
	q := CBTI.Queue
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}

		CBTI.Queue = append(CBTI.Queue, q[0])

		q = q[1:]
	}
	return CBTI
}

//进行插入操作，树为完全二叉树（插满一层再到下一层）
func (this *CBTInserter) Insert(val int) int {
	var lastNode, node *TreeNode
	node = &TreeNode{Val: val}
	for i := 0; i < len(this.Queue); i++ {
		if this.Queue[i].Left == nil {
			lastNode = this.Queue[i]
			lastNode.Left = node
			break
		}
		if this.Queue[i].Right == nil {
			lastNode = this.Queue[i]
			lastNode.Right = node
			break
		}
	}
	this.Queue = append(this.Queue, node)
	return lastNode.Val
}

//返回完全二叉树的根节点
func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}