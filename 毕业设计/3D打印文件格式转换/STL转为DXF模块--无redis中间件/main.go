package main

import (
	"2022.9Project/StlToStruct/StlToStruct"
	"2022.9Project/StlToStruct/StructToDXF"
	"fmt"
	"os"
)

func main() {
	var err error
	fileObj, err := os.Open("/home/keqing/桌面/芭芭拉.stl")
	if err != nil {
		fmt.Println("文件打开失败！")
		return
	}
	defer fileObj.Close()

	//modelInfo := make([]byte, 80)
	//fileObj.Read(modelInfo)
	_, err = fileObj.Seek(80, 0)
	if err != nil {
		fmt.Println("文件指针移动失败!")
		return
	}

	num := make([]byte, 4)
	_, err = fileObj.Read(num)
	if err != nil {
		fmt.Println("三角形数量读取失败！")
		return
	}

	UnitList := StlToStruct.ReadFile(fileObj, num)
	StructToDXF.StructToDXF(UnitList)
}
