package StlToStruct

import (
	"bufio"
	"encoding/binary"
	"math"
	"os"
)

type AngleUnit struct {
	VertexA []float32
	VertexB []float32
	VertexC []float32
}

func cal(idx int) int {
	return 1 << (8 * idx)
}
//字节转int数据
func ByteToInt(data []byte) int {
	var res int
	for i, val := range data {
		res += cal(i) * int(val)
	}
	return res
}
//字节转float数据
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}
//获取面单元
func GetUnit(data []byte) *AngleUnit {
	Unit := &AngleUnit{}
	for i := 4; i <= 48; i += 4 {
		if i > 12 && i <= 24 {
			Unit.VertexA = append(Unit.VertexA, ByteToFloat32(data[i-4:i]))
		}
		if i > 24 && i <= 36 {
			Unit.VertexB = append(Unit.VertexB, ByteToFloat32(data[i-4:i]))
		}
		if i > 36 && i <= 48 {
			Unit.VertexC = append(Unit.VertexC, ByteToFloat32(data[i-4:i]))
		}
	}
	return Unit
}

func ReadFile(fileObj *os.File, numByte []byte) []*AngleUnit {
	UnitList := make([]*AngleUnit, 0, ByteToInt(numByte))
	//自定义缓冲
	reader := bufio.NewReaderSize(fileObj, 50*1024*1024)
	//读取面单元数据
	for {
		BufData := make([]byte, 50)
		n, _ := reader.Read(BufData)
		if n == 0 {
			break
		}
		UnitList = append(UnitList, GetUnit(BufData))
	}
	return UnitList
}
