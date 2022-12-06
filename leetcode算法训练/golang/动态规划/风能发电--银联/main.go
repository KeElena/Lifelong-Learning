func storedEnergy(storeLimit int, power []int, supply [][]int) int {
    
    order:=make(map[int][]int,len(supply))      //将指令转为map进行存储
    var ord int                                 //定义执行的指令
    var eng int                                 //定义储能站剩余能源
    
    for _,s:=range supply{                      //遍历supply，将二维数组转为散列表
        order[s[0]]=s[1:]
    }
    
    for i,elect:=range power{                   //遍历能源发电情况
        if order[i]!=nil{                       //查看当前时刻是否有变更指令
            ord=i                               //有则更换指令
        }

        if elect<order[ord][0]{                   //判断实时发电量是否小于最小供应电能
            eng-=order[ord][0]-elect              //是则消耗储能站的资源
            if eng<0{                             //储能站默认不能低于0
                eng=0
            }
        }
        
        if elect>order[ord][1]{                   //判断实时发电量是否大于最大供应电能
            eng+=elect-order[ord][1]              //是则增加储能站的资源
            if eng>storeLimit{                    //储能站的资源不能超过storeLimit
                eng=storeLimit
            }
        }
    }
    return eng                                    //返回储能站资源情况
}