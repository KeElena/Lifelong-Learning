package main

import "fmt"

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func buildTree(root *treeNode) *treeNode {
	var val int
	fmt.Scanf("%d\n", &val)
	if val == 3 {
		return nil
	} else {
		root.Val = val
		root.Left = buildTree(&treeNode{})
		root.Right = buildTree(&treeNode{})
	}
	return root
}

func pruneTree(root *treeNode) *treeNode {
	if root == nil {
		return nil
	} else {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
		if root.Left == nil && root.Right == nil && root.Val == 0 {
			return nil
		}
	}
	return root
}

func main() {

	//root := &treeNode{Val: 1}
	//root.Right = &treeNode{Val: 0}
	//root.Right.Left = &treeNode{Val: 0}
	//root.Right.Right = &treeNode{Val: 1}

	root := buildTree(&treeNode{})
	pruneTree(root)
	fmt.Println(root)
}
