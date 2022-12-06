package main

import (
	"fmt"
)

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

type Node struct {										//定义三元组结构
	x     int											//节点的值
	left  int											//左编号
	right int											//右编号
}

func isExist(node *Node, tree []*Node) (bool, int) {	//判断树节点是否重复
	for i, n := range tree {							//遍历树
		if *node == *n {								//结构体的值相同时，返回切片内相同元素的索引和true
			return true, i								//结构体是值类型，相同类型的结构体可以直接进行比较
		}
	}
	return false, -1									//否则返回false和-1（切片不存在-1的索引）
}

func do(root *TreeNode) []*TreeNode {
	tree := make([]*Node, 0, 5)							//存储大量数据且索引为数字时推荐用切片，散列表的构造和存取非常消耗时间
	store := make(map[int]*TreeNode)					//去除重复数据推荐用散列表，一般用于存储结果数据，数量相对少

	var i int											//声明三元组编号的索引
	var n *Node											//声明Node结构体指针
	var dfs func(root *TreeNode) int					//声明闭包递归函数

	dfs = func(root *TreeNode) int {					//构造闭包递归函数，使用后序遍历
		if root == nil {								//节点为空时
			return -201									//返回一个范围之外的默认为空的数
		}

		left := dfs(root.Left)							//左递归
		right := dfs(root.Right)						//右递归

		n = &Node{x: root.Val, left: left, right: right}//构造三元组
		ok, index := isExist(n, tree)					//接收返回的参数
		if ok {											//如果存在重复
			store[index] = root							//根据索引值存储树节点
			return index								//（编号是前面的，物理树节点是后面的，物理节点相同所以不影响结构）
		}												//返回三元组编号

		tree = append(tree, n)							//没有重复则追加树节点
		i++												//编号+1
		return i - 1									//返回新增的三元组的编号
	}
	dfs(root)											//启动递归

	res := make([]*TreeNode, 0, len(store))				//构造结构集
	for _, node := range store {						//遍历store散列表
		res = append(res, node)							//追加树节点
	}
	return res											//返回结果
}

func main() {
	arr := []int{1, 2, 4, 0, 0, 0, 3, 2, 4, 0, 0, 0, 4, 0, 0}
	root, _ := buildTree(&TreeNode{}, arr)
	fmt.Println(do(root))
}
