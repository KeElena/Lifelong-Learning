package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(root *TreeNode, arr []int) (*TreeNode, []int) {

	if arr[0] == 111 {
		return nil, arr[1:]
	} else {
		root.Val = arr[0]
		arr = arr[1:]
		root.Left, arr = buildTree(&TreeNode{}, arr)
		root.Right, arr = buildTree(&TreeNode{}, arr)
	}
	return root, arr
}

func do(root *TreeNode) int {

	id := make(map[int][]int, 3)							//构造散列表
	max := 0												//最大宽度

	var DFS func(root *TreeNode, rootNum int, layer int)	//声明闭包递归函数
	DFS = func(root *TreeNode, rootNum int, layer int) {	//定义闭包递归函数
		if root == nil {									//节点为空时返回
			return
		}

		id[layer] = append(id[layer], rootNum)				//根据层数进行插值

		DFS(root.Left, 2*rootNum, layer+1)					//左递归，满二叉树的情况下左节点的编号为 2*rootNum
		DFS(root.Right, 2*rootNum+1, layer+1)				//右递归，满二叉树的情况下右节点的编号为 2*rootNum+1

		return
	}
	DFS(root, 1, 0)											//执行递归函数

	for _, s := range id {									//遍历散列表的切片
		if s[len(s)-1]-s[0]+1 > max {						//获取最大宽度
			max = s[len(s)-1] - s[0] + 1
		}
	}

	return max												//返回最大宽度
}

func main() {
	//arr := []int{1, 3, 5, 111, 111, 3, 111, 111, 2, 111, 9, 111, 111}
	arr := []int{1, 3, 5, 6, 111, 111, 111, 111, 2, 111, 9, 7, 111, 111, 111}
	//arr := []int{1, 3, 5, 111, 111, 111, 2, 111, 111}
	root, _ := buildTree(&TreeNode{}, arr)
	fmt.Println(do(root))
}
