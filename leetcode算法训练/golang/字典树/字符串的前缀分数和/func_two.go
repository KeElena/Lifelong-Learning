//字典树节点结构
type charNode struct {
	next  [26]*charNode							//next指针有26种可能（26个字母）
	score int									//节点访问次数
}
//构造字典树和记录节点访问次数
func dfs(word string, i int, root *charNode) {
	root.score++								//节点访问次数+1

	if i == len(word) {							//一条路径有len(word)+1个节点（开始的节点为字典树的root，不属于任何字符串）
		return									//遍历到字符串最后一个元素时返回
	}

	if root.next[word[i]-97] == nil {			//如果下一个字符的节点为空，则构造一个新的节点
		root.next[word[i]-97] = &charNode{}
	}

	dfs(word, i+1, root.next[word[i]-97])		//递归构造字典树

	return										//返回
}

func count(root *charNode, word string, i int) int {//递归的方式获取路径上节点的访问次数
	if i == len(word)-1 {							//递归到最后一个字符时只返回score
		return root.score
	}
	return root.score + count(root.next[word[i+1]-97], word, i+1)//递归累加节点的访问次数
}

func do(words []string) []int {
    res := make([]int,len(words))								//构造结果集

	root := &charNode{}											//构造字典树的根节点

	for _, word := range words {								//遍历每个字符串并构造出字典树
		dfs(word, 0, root)
	}

	for i, word := range words {								//统计每个字符串路径的计数
		res[i]+=count(root.next[word[0]-97], word, 0)			//输出计数
	}
	return res													//返回结果
}