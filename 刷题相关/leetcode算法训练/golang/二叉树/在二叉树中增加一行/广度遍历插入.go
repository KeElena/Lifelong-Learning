package main

import "fmt"
//树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//构造树
func buildTree(root *TreeNode, arr []int) (*TreeNode, []int) {

	if arr[0] == 0 {
		return nil, arr[1:]
	} else {
		root.Val = arr[0]
		arr = arr[1:]
		root.Left, arr = buildTree(&TreeNode{}, arr)
		root.Right, arr = buildTree(&TreeNode{}, arr)
	}
	return root, arr
}
//遍历树
func traverseTree(node *TreeNode) {
	if node == nil {
		return
	} else {
		fmt.Printf("%d ", node.Val)
		traverseTree(node.Left)
		traverseTree(node.Right)
	}
	return
}

func do(root *TreeNode, val int, depth int) *TreeNode {

	if depth == 1 {								//深度为1时直接处理
		return &TreeNode{Val: val, Left: root}
	}

	Queue := make([]*TreeNode, 0, 5)			//创建树枝队列
	Queue = append(Queue, root)					//增加根节点
	queueSize := 0								//初始化队列大小

	for len(Queue) > 0 {

		if depth == 2 {							//深度为2时
			for _, node := range Queue {		//循环为层的使用节点进行插入操作
				node.Left = &TreeNode{Val: val, Left: node.Left}
				node.Right = &TreeNode{Val: val, Right: node.Right}
			}
			return root							//完成插入操作后返回根节点并结束函数
		}

		queueSize = len(Queue)					//获取层节点的个数
		for _, node := range Queue {			//循环遍历层节点追加下一层的节点
			if node.Left != nil {
				Queue = append(Queue, node.Left)
			}
			if node.Right != nil {
				Queue = append(Queue, node.Right)
			}
		}

		Queue = Queue[queueSize:]				//裁剪当层的节点
		depth--									//深度-1
	}
	return nil
}

func main() {
	arr := []int{5, 3, 0, 0, 1, 4, 0, 2, 0, 0, 0}
	root, _ := buildTree(&TreeNode{}, arr)
	traverseTree(do(root, 2, 4))
}
