package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {

	//计算数组长度
	line,row := len(grid),len(grid[0])
	s := make([]int, 0, line*row)
	//转为一维数组
	for _, x := range grid {
		for _, y := range x {
			s = append(s, y)
		}
	}
	//数组重构
	for i := 0; i < len(s); i++ {
		addr := (i + k) % (line * row)
		grid[addr/row][addr%row] = s[i]
		//y := addr % row
		//x := addr / row
	}
	//返回重构后的数组
	return grid
}

func main() {
	arr := [][]int{
		{1, 2, 3},
		{3, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(shiftGrid(arr, 23))

}
