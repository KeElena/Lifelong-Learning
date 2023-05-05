func dfs(root *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}

	root.Left=dfs(root.Left)
	root.Right=dfs(root.Right)
    
	return root
}

