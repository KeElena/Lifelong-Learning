func numSpecial(mat [][]int) int {
    
    rows:=make([]int,0,len(mat))        //存储行和为1的行，len(mat)表示行数
    cols:=make([]int,0,len(mat[0]))     //存储列和为1的列，len(mat[0])表示列数
    var sum int                         //定义累加参数

    for index,arr:=range mat{           //按行遍历二维数组
        sum=0                           //sum归0
        for _,val:=range arr{           //元素值累加
            sum+=val
        }
        if sum==1{                      //累加完后和为1则存储该行的索引
            rows=append(rows,index)
        }
    }

    for x:=0;x<len(mat[0]);x++{         //按列遍历二维数组
        sum=0                           //sum归0
        for y:=0;y<len(mat);y++{        //元素值累加
            sum+=mat[y][x]
        }
        if sum==1{                      //累加完后和为1时存储该列
            cols=append(cols,x)
        }
    }
    sum=0                               //累加元素归0
    for _,row:=range rows{              //遍历rows和cols数组
        for _,col:=range cols{
            if mat[row][col]==1{        //判断交叉位置的值是否为1
                sum++                   //是时结果累计+1
            }
        }
    }
    return sum                          //返回结果
}
//特殊位置的特性：行与列累加和均为1
//行与列累加和为1，但交叉位置的值不一定是1
//需要获取累加为1的行和列，然后判断交叉位置是否为1