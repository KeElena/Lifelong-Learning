package ParseSTL

import (
	"encoding/binary"
	"math"
)

type AngleUnit struct {
	VertexA []float32
	VertexB []float32
	VertexC []float32
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

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
