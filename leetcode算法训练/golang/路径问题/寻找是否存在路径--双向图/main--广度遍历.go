//使用广度遍历寻找是否有路径到达目的地
func validPath(n int, edges [][]int, source int, destination int) bool {
    //特殊情况，起点与目的相同
    if source==destination{
        return true
    }
    //使用队列缓存，用于暂存同层的顶点
    cache:=make([]int,0,n)
    //起点加入队列缓存
    cache=append(cache,source)
    //使用状态表，用于记录顶点是否被搜索过
    status:=make([]bool,n)
    size:=0
    //左顶点id 右顶点id，双向图时
    var left,right int
    //缓存队列有值时循环
    for len(cache)>0{
        //记录缓存队列的大小
        size=len(cache)
        //遍历缓存队列里的所有顶点
        for _,v:=range cache{
            //遍历所有边
            for i:=0;i<len(edges);i++{
                //记录左顶点
	left=edges[i][0]
                //记录右顶点
                right=edges[i][1]
                //左顶点为缓存内的顶点且右顶点没有被搜索过时
                if left==v && status[right]==false{
	    //如果是目的地则返回true
                    if right ==destination{
                        return true
                    }
	    //顶点放入缓存队列
                    cache=append(cache,right)
                //右顶点为缓存内的顶点且左顶点没有被搜索过时
                }else if right==v && status[left]==false{
	    //如果是目的地则返回true
                    if left ==destination{
                        return true
                    }
	    //顶点放入缓存队列
                    cache=append(cache,left)
                }
            }
            //将搜索过的顶点记录为true
            status[v]=true
        }
        //去掉搜索过的点
        cache=cache[size:]
    }
    return false
}