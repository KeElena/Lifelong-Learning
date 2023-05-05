//使用并查集处理，如果起点和目的的祖先相同，则说明两点间一定存在一条路径
func validPath(n int, edges [][]int, source int, destination int) bool {
    set:=initSet(n)
    for _,edge:=range edges{
        union(set,edge[0],edge[1])
    }
    fmt.Println(find(set,source),find(set,destination))
    if find(set,source)==find(set,destination){
        return true
    }
    return false
}

func initSet(n int)(arr []int){
    arr=make([]int,n)
    for i:=range arr{
        arr[i]=i
    }
    return
}

func find(set []int,p int)int{
    if set[p]==p{
        return p
    }
    
    set[p]=find(set,set[p])
    return set[p]
}

func union(set[]int,left int,right int){
    if left<right{
       set[find(set,left)]=find(set,right)
    }else{
        set[find(set,right)]=find(set,left)
    }
}