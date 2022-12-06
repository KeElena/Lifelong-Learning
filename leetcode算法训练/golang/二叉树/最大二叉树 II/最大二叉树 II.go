func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {

    if val>root.Val{                                            //插入值大于根的值
        root=&TreeNode{Val:val,Left:root}                       //换根
    }else{
        var cur,precur *TreeNode                                //定义当前节点地址和上一个节点地址
        cur=root                                                //初始化当前节点为root
        for {                                                   //循环
            if cur.Right!=nil && cur.Val>val{                   //下一个右节点非空且当前节点的值大于val时执行
                precur=cur                                      //记录当前节点地址
                cur=cur.Right                                   //记录下一个右子树节点地址
            }else{
                break                                           //不满足条件则退出循环
            }
        }

        if val<cur.Val{                                         //val小于cur树节点的值，构造的树节点插入到cur节点的右子树上
            cur.Right=&TreeNode{Val:val}
        }else{                                                  //val大于cur树节点的值时，构造的树节点需要插入到cur节点的前面
            precur.Right=&TreeNode{Val:val,Left:precur.Right}   //使用上一个节点完成插入构造节点，将cur节点插到构造节点的左子树上
        }
    }

    return root
}