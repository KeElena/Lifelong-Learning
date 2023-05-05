package func_two

type TreeNode struct{
	Val 	int
	Left 	*TreeNode
	Right 	*TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {										//遇到空节点则返回nil
		return nil
	}

	if depth == 1 {											//深度为1处理完后返回对应的根节点
		return &TreeNode{Val: val, Left: root}
	}
	if depth == 2 {
		root.Left = &TreeNode{Val: val, Left: root.Left}	//深度为2时进行节点的插入操作
		root.Right = &TreeNode{Val: val, Right: root.Right}	
	} else {												//完成插入操作后不需要递归，使用else跳过递归
		root.Left = addOneRow(root.Left, val, depth-1)		//未达到深度时进行递归
		root.Right = addOneRow(root.Right, val, depth-1)
	}

	return root												//完成插入操作或递归返回时返回根节点
}
