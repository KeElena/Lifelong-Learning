func setZeroes(matrix [][]int)  {
    line:=make([]int,len(matrix))       //使用切片去冗余标记line
    row:=make([]int,len(matrix[0]))     //使用切片去冗余标记row

    for i,arr:=range matrix{
        for j,val:=range arr{           //满足val=0时记录行和列
            if val==0{
                line[i]=1
                row[j]=1
            }
        } 
    }
    for i:=0;i<len(line);i++{           //遍历行元素
        if line[i]==1{                  //line[i]==1表示i为标记的行
            LineToZero(matrix,i)
        }
    }
    for i:=0;i<len(row);i++{            //row[i]==1表示i为标记的列
        if row[i]==1{
            RowToZero(matrix,i)
        }
    }
}
func LineToZero(matrix [][]int,line int){   //将行所有元素转为0
    for i:=0;i<len(matrix[0]);i++{
        matrix[line][i]=0
    }
}
func RowToZero(matrix [][]int,row int){     //将列所有元素转为0
    for i:=0;i<len(matrix);i++{
        matrix[i][row]=0
    }
}
