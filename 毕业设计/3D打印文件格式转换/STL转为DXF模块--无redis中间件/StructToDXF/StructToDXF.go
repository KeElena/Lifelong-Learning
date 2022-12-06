package StructToDXF

import (
	unit "2022.9Project/StlToStruct/StlToStruct"
	"fmt"
	"os"
)

func StructToDXF(UnitList []*unit.AngleUnit) {
	//获取文件对象
	f, err := os.OpenFile("./test.dxf", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//初始化文件内容
	content := make([]byte, 0, 5*1024*1024)
	//写入dxf实体单元开头
	content = append(content, []byte("0\nSECTION\n2\nENTITIES\n")...)
	//写入dxf实体单元数据
	for _, face := range UnitList {
		content = append(content, []byte("0\n3DFACE\n8\n0\n")...)
		content = append(content, []byte(fmt.Sprintf("10\n%s\n20\n%s\n30\n%s\n", face.VertexA[0], face.VertexA[1], face.VertexA[2]))...)
		content = append(content, []byte(fmt.Sprintf("11\n%s\n21\n%s\n31\n%s\n", face.VertexB[0], face.VertexB[1], face.VertexB[2]))...)
		content = append(content, []byte(fmt.Sprintf("12\n%s\n22\n%s\n32\n%s\n", face.VertexC[0], face.VertexC[1], face.VertexC[2]))...)
		content = append(content, []byte(fmt.Sprintf("13\n%s\n23\n%s\n33\n%s\n", face.VertexC[0], face.VertexC[1], face.VertexC[2]))...)
	}
	//写入文件结尾
	content = append(content, []byte("0\nENDSEC\n0\nEOF\n")...)
	//存储内容
	f.Write(content)
}
