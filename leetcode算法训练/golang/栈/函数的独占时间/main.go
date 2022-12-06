package main

import (
	"strconv"
	"strings"
)

type elem struct {								//栈元素的结构
	Start    int								//开始位置
	FuncName int								//方法名
	Stop     int								//断点
}

func do(n int, logs []string) []int {

	stack := make([]elem, 0, n)					//创建一个栈切片
	res := make([]int, n, n)					//创建输出结果队列
	var i int									//声明一个栈顶索引值

	for _, str := range logs {					//遍历输入的日志
		funcNum, _ := strconv.Atoi(str[:strings.Index(str, ":")])		//获取函数名
		num, _ := strconv.Atoi(str[strings.LastIndex(str, ":")+1:])		//获取时间轴的点
		i = len(stack) - 1												//获取栈顶索引
		//占用时间追加公式：x+=end-start/stop+1(右边计算增加的时间大小)
		if strings.Contains(str, "start") {								//如果操作为start，进行入栈相关操作
			if len(stack) > 0 {											//栈元素个数大于0时
				if stack[i].Stop != 0 {									//元素断点非0时
					res[stack[i].FuncName] += num - stack[i].Stop		//根据断点追加占用时间
				} else {												//断点为0，即无断点
					res[stack[i].FuncName] += num - stack[i].Start		//无断点时根据start值追加占用时间
				}
				stack[i].Stop = num										//更新断点
			}
			stack = append(stack, elem{Start: num, FuncName: funcNum})	//进行入栈
		}

		if strings.Contains(str, "end") {								//如果操作为end，进行出栈相关操作
			if stack[i].Stop != 0 {										//操作断点时
				res[stack[i].FuncName] += num - stack[i].Stop + 1		//根据断点进行追加占用时间
			} else {
				res[stack[i].FuncName] += num - stack[i].Start + 1		//没有断点时根据start追加占用时间
			}

			if len(stack) > 1 {											//栈元素大于1时更新栈元素上一个元素的断点
				stack[i-1].Stop = num + 1
			}
			stack = stack[:i]											//出栈,左闭右开，将第i个元素出栈
		}
	}
	return res															//返回结果
}

func main() {
	n := 2
	logs := []string{"0:start:0", "0:end:0"}
	//logs := []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}
	do(n, logs)
}
