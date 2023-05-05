func largestLocal(grid [][]int) [][]int {
    res:=make([][]int,len(grid)-2)
    for i:=0;i<len(grid)-2;i++{
        for j:=0;j<len(grid)-2;j++{
            res[i]=append(res[i],getMax(i,j,grid))
        }
    }
    return res
}
func getMax(x,y int,grid [][]int)int{
    var max int
    for i:=x;i<x+3;i++{
        for j:=y;j<y+3;j++{
            if grid[i][j]>max{
                max=grid[i][j]
            }
        }
    }
    return max
}