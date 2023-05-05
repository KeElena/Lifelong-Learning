package main

import "fmt"
//树节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//用于构造树的递归函数
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

func printTree(root *TreeNode) [][]string {
	
	var deep int
	//使用闭包定义递归函数
		//声明
	var getDeep func(root *TreeNode, num int)				//闭包构造递归函数必须要声明，否则不能递归
		//构造
	getDeep = func(root *TreeNode, num int) {
		if root == nil {
			if num > deep {
				deep = num
			}
			return
		}
		getDeep(root.Left, num+1)
		getDeep(root.Right, num+1)
		return
	}
	//获取最大深度
	getDeep(root, 0)
	//通过深度计算矩阵的宽度
	col := sqr(deep) - 1
	//初始化矩阵
	res := make([][]string, 0, deep)
	for i := 0; i < deep; i++ {
		res = append(res, make([]string, col))
	}
	//使用闭包构造递归式的转换函数
		//声明递归函数
	var transform func(root *TreeNode, setAddr int, layer int)
		//构造递归函数
	transform = func(root *TreeNode, setAddr int, layer int) {
		//节点为空时返回
		if root == nil {								//处理
			return
		}
		//指定位置插入值
		res[layer][setAddr] = fmt.Sprint(root.Val)
		//左右指针为空时返回
		if root.Left == nil && root.Right == nil {		//处理完叶子节点则返回，不用递归下去，否则导致移位运算错误
			return
		}
														//deep-laye-2一定要大于等于0
		transform(root.Left, setAddr-1<<(deep-layer-2), layer+1)
		transform(root.Right, setAddr+1<<(deep-layer-2), layer+1)
		return
	}
	//执行递归函数
	transform(root, (col-1)/2, 0)

	return res
}
//通过移位运算计算平方
func sqr(n int) (res int) {
	return 1 << n
}

func main() {
	arr := []int{1, 3, 3}
	root, _ := buildTree(&TreeNode{}, arr)
	fmt.Println(printTree(root))
}
