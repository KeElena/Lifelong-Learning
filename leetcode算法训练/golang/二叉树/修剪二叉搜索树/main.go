func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root==nil{
        return nil
    }

    if root.Val>=low{                           //根节点的值大于等于low时右递归
        root.Left=trimBST(root.Left,low,high)
    }
    if root.Val<=high{                          //根节点的值小于等于high时左递归
        root.Right=trimBST(root.Right,low,high)
    }

    if root.Val<low{                            //根节点的值小于low时返回右节点
        return root.Right
    }
    if root.Val>high{                           //根节点的值大于hight时返回左节点
        return root.Left
    }
    return root
}