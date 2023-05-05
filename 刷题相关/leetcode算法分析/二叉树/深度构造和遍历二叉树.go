package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(root *TreeNode, arr []int) (*TreeNode, []int) {

	if arr[0] == 3 {
		return nil, arr[1:]
	} else {
		root.Val = arr[0]
		arr = arr[1:]
		root.Left, arr = buildTree(&TreeNode{}, arr)
		root.Right, arr = buildTree(&TreeNode{}, arr)
	}
	return root, arr
}

func traverseTree(root *TreeNode) {
	if root == nil {
		return
	} else {
		fmt.Printf("%d ", root.Val)
		traverseTree(root.Left)
		traverseTree(root.Right)
	}
	return
}
