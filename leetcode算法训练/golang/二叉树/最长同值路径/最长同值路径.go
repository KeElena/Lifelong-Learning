package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func longestUnivaluePath(root *TreeNode) int {
    //定义递归函数
    var dfs func(*TreeNode) int
    //定义使用到的参数
    var ans,num1,num2 int           //定义num1和num2，进行闭包递归时参数指向的地址都是原来的地址
                                    //不要用闭包环境的变量接收闭包递归函数的值，否则值会频繁改变导致bug
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        //使用后序遍历
        left := dfs(node.Left)
        right := dfs(node.Right)
        //参数置0
        num1,num2=0,0
        //左右节点与父节点的值相同时，继承路径的边数
        if node.Left != nil && node.Left.Val == node.Val {
            num1=left+1
        }
        if node.Right != nil && node.Right.Val == node.Val {
            num2=right+1
        }
        //结果选取路径最大值
        ans = max(ans, num1+num2)
        //左子树和右子树中选取最大的子路径然后返回（路径左右只有两个端点）
        return max(num1,num2)
    }
    dfs(root)
    return ans
}
//如果根节点和左节点值不相等，那么从该根节点起到左边的链路最大就是0，所以要新设一个变量
//left直接++，永远都是默认当前根节点和左右值相等，左右递归结果相加。

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//arr := []int{1, 2, 2, 2, 0, 0, 0, 2, 0, 0, 2, 2, 0, 0, 2, 0, 0}
	//arr := []int{2, 2, 2, 2, 0, 0, 0, 2, 0, 0, 0}
	arr := []int{1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0}
	root, _ := buildTree(&TreeNode{}, arr)
	fmt.Println(longestUnivaluePath(root))
}
