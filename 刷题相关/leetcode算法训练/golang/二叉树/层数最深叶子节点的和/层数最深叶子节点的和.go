 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func deepestLeavesSum(root *TreeNode) int {
    queue:=make([]*TreeNode,0,3)                    //构造树节点队列
    queue=append(queue,root)                        //添加树根
    bol,size,sum:=false,0,0                         //需要的参数

    for len(queue)>0{                               //队列长度大于0时循环
        size=len(queue)                             //获取层节点数量
        for _,node:=range queue{                    //循环层节点
            if node.Left !=nil{                     //左指针非空时追加节点
                queue=append(queue,node.Left)
                bol=true                            //标志置为true
            }
            if node.Right !=nil{                    //右指针非空时追加节点
                queue=append(queue,node.Right)
                bol=true                            //标志置位true
            }
        }

        if bol == false{                            //如果没追加节点标志为false，即为最底层
            fmt.Println(queue)        
            for _,node:=range queue{                //累加运算
                sum+=node.Val
            }
            return sum                              //返回结果
        }

        queue=queue[size:]                          //裁剪队列
        bol=false                                   //归零
    }
    return 0
}