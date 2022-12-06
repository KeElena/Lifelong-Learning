package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {

	Queue :=make([]*TreeNode,0,5)								//初始化
    Queue=append(Queue,root)									//添加根节点
	size, layer, minLayer, maxSum, sum := 1, 1, 0, -10001, 0	//初始化参数

	for len(Queue) > 0 {										//循环遍历

	    size = len(Queue)        								//获取队列长度

		for _, node := range Queue {							//for-range循环，添加元素不会影响循环次数
			sum += node.Val										//求和
			if node.Left != nil {								//添加左节点
				Queue = append(Queue, node.Left)
			}
			if node.Right != nil {								//添加右节点
				Queue = append(Queue, node.Right)
			}
		}

		if sum > maxSum {										//求是否是最大的求和
			minLayer = layer									//是则写入层数
			maxSum = sum										//保存最大求和的值
		}

        sum-=sum												//sum清0
		Queue = Queue[size:]									//根据size裁剪队列
		layer++													//层数+1
	}
	return minLayer												//返回求和最大值的最小层
}
