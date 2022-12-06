//初始化并查集
func initArr(n int) []int {
	set := make([]int, n)											//初始化集合
	for i := range set {											//初始化每个元素的头节点索引
		set[i] = i													//初始化默认头节点为本身索引
	}
	return set
}
//使用路径压缩查找元素i的头节点
func find(i int, set []int) int {

	if set[i] == i {												//如果头节点为本身，则返回元素索引
		return i
	}
	set[i] = find(set[i], set)										//使用递归接收传回的头节点，实现路径压缩
	return set[i]													//返回接收的头节点
}
//将i合并到j的集合里
func union(i, j int, set []int) {
	i_Fa := find(i, set)											//查找i的祖先
	j_Fa := find(j, set)											//查找j的祖先
	set[i_Fa] = j_Fa												//更换i的祖先的头结点为j的祖先
}