func largestIsland(grid [][]int) int {

	var max, sum, temp int

	set := initArr(len(grid) * len(grid))						//初始化并查集合
	count := Area(grid, set)									//获取每个大陆的面积

	mem := make([]int, 4)										//创建存储与坐标相连的大陆
	for line, arr := range grid {
		for row, val := range arr {								//遍历每个坐标
			if val == 1 {										//val为1时跳过
				continue
			}

			sum = 1												//陆地面积默认为1（1为海洋变为大陆的面积）
			mem = mem[:0]
			if line+1 < len(grid) && grid[line+1][row] == 1 { 	//向下查询是否有大陆
				temp = find((line+1)*len(arr)+row, set)			//有则获取大陆的祖先
				if valueNoIn(mem, temp) {						//检测是否重复
					mem = append(mem, temp)						//存储相连的大陆
				}
			}
			if line-1 > -1 && grid[line-1][row] == 1 { 			//向上查询是否有大陆
				temp = find((line-1)*len(arr)+row, set)			//有则获取大陆的祖先
				if valueNoIn(mem, temp) {						//检测是否重复
					mem = append(mem, temp)						//存储相连的大陆
				}
			}
			if row+1 < len(grid[0]) && grid[line][row+1] == 1 { //向右查询是否有大陆
				temp = find(line*len(arr)+row+1, set)			//有则获取大陆的祖先
				if valueNoIn(mem, temp) {						//检测是否重复
					mem = append(mem, temp)						//存储相连的大陆
				}
			}
			if row-1 > -1 && grid[line][row-1] == 1 { 			//向左查询是否有大陆
				temp = find(find(line*len(arr)+row-1, set), set)//有则获取大陆的祖先
				if valueNoIn(mem, temp) {						//检测是否重复
					mem = append(mem, temp)						//存储相连的大陆
				}
			}
			for _, v := range mem {								//遍历所有相连的大陆
				sum += count[v]									//获取对应大陆的面积并相加
			}
			max = Max(max, sum)									//在所有与大陆相连的坐标中选取最大的面积
		}
	}
	if max == 0 {												//max为0表示全为大陆
		return len(grid) * len(grid)							//直接返回整个大陆的面积
	}
	return max
}

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
//获取每个岛的面积
func Area(grid [][]int, set []int) map[int]int {
	count := make(map[int]int, 5)									//使用map表进行计数
	for line, arr := range grid {
		for row, val := range arr {									//遍历每个坐标
			if val == 0 {											//将非大陆坐标的头结点置为-1，表示剔除
				set[line*len(arr)+row] = -1
				continue
			}

			if line+1 < len(grid) && grid[line+1][row] == 1 { 		//检测大陆并向下合并陆地集合
				union(line*len(arr)+row, (line+1)*len(arr)+row, set)
			}
			if row+1 < len(grid[0]) && grid[line][row+1] == 1 { 	//检测大陆并向右合并陆地集合
				union(line*len(arr)+row, line*len(arr)+row+1, set)
			}
			if row-1 > -1 && grid[line][row-1] == 1 { 				//检测大陆并向左合并陆地集合
				union(line*len(arr)+row, line*len(arr)+row-1, set)
			}
		}
	}

	for _, val := range set {										//遍历set集合
		if val == -1 {
			continue
		}
		count[find(val, set)]++										//统计所有大陆的面积
	}
	return count													//返回map表
}
//获取最大值
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
//检测是否重复
func valueNoIn(arr []int, val int) bool {
	for _, num := range arr {
		if num == val {
			return false
		}
	}
	return true
}