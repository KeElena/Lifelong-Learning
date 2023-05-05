package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type construct struct {
	Root  *TreeNode
	Queue []*TreeNode
}

func getInput(constructor *construct) *construct {
	for i := 0; i < 4; i++ {
		constructor.Queue = append(constructor.Queue, &TreeNode{Val: i})
	}
	return constructor
}

func buildTree(constructor *construct) *TreeNode {

	i := 1
	q := constructor.Queue
	for i < len(constructor.Queue) {

		if q[0].Left == nil {
			q[0].Left = constructor.Queue[i]
			i++
		}
		if i == len(constructor.Queue) {
			break
		}
		if q[0].Right == nil {

			q[0].Right = constructor.Queue[i]
			i++
		}
		q = q[1:]
	}
	return constructor.Queue[0]
}

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

//进行广度遍历并存储地址
func Constructor(root *TreeNode) construct {

	var cache = construct{Root: root}
	q := cache.Queue
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}

		cache.Queue = append(cache.Queue, q[0])
		q = q[1:]
	}
	return cache
}

func main() {

	constructor := getInput(&construct{})
	root := buildTree(constructor)
	traverseTree(root)

}
