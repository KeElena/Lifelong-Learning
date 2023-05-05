package main

import "fmt"

func do(grid [][]int) int {

	spot := make([][]int, 0, len(grid))								//存储与两个或以上的岛相连的位置
	var sum int
	var max int
	var pc int														//记录地图总面积

	for line, arr := range grid {									
		for row, val := range arr {									//遍历每个位置
			if val == 1 {
				pc++												//坐标有陆地时总面积+1
				continue											//跳过循环
			}
			sum = 0													//重置计数器
			if line+1 < len(grid) && grid[line+1][row] == 1 { 		//检测坐标下边有无陆地
				sum++
			}
			if line-1 > -1 && grid[line-1][row] == 1 {				//检测坐标上边有无陆地
				sum++
			}
			if row+1 < len(grid[0]) && grid[line][row+1] == 1 { 	//检测坐标右边有无陆地
				sum++
			}
			if row-1 > -1 && grid[line][row-1] == 1 { 				//检测坐标左边有无陆地
				sum++
			}
			if sum > 1 {											//计数大于1时存储坐标
				spot = append(spot, []int{line, row})
			}
		}
	}

	set := initArr(len(grid) * len(grid))							//初始化并查集合

	if len(spot) != 0 {												//如果存在坐标与两个或以上的大陆相连
		for _, p := range spot {									//遍历每个坐标
			max = Max(max, getMaxArea(grid, p, set))				//输入坐标获取最大区域的值，比较之前的最大值选取最大的那个
		}
	} else {														//如果不存在与两个或以上的大陆相连的坐标
		if pc == 0 {												//大陆面积为0时返回1
			return 1
		} else if pc == len(grid)*len(grid[0]) {					//大陆面积为最大面积时返回最大面积
			return pc
		}
		max = getMaxArea(grid, nil, set)							//不输入坐标返回最大的面积
		max++														//将一个海洋坐标换成大陆
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
//获取最大区域的面积
func getMaxArea(grid [][]int, spot []int, set []int) (max int) {
	defer func() {													//函数运行完重置set集合
		for i := range set {
			set[i] = i
		}
	}()

	if spot != nil {												//如果坐标非空
		grid[spot[0]][spot[1]] = 1									//将对应坐标变为陆地
		defer func() { grid[spot[0]][spot[1]] = 0 }()				//程序结束后恢复为海洋
	}

	for line, arr := range grid {
		for row, val := range arr {									//遍历所有坐标
			if val == 0 {											//坐标为0时，将对应set的元素的头结点置为-1，意为剔除
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
	max = getMax(set)												//获取面积最大的大陆
	return
}

func getMax(set []int) int {
	count := make(map[int]int, 5)									//构造一个map表用于计数

	var max int
	for _, val := range set {										//遍历查并集合的元素
		if val == -1 {												//val为-1时跳过
			continue
		}
		count[find(val, set)]++										//查找元素的祖先，对应祖先的计数+1
	}

	for _, val := range count {										//遍历map表，返回最大的陆地
		max = Max(max, val)
	}
	return max
}
//获取最大值
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	grid := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{1, 0, 0}}
	//[[0,1,0],[1,0,1],[1,0,0]]
	fmt.Println(do(grid))
}
