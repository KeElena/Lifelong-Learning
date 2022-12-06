package GetFile

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"strings"
)

type AngleUnit struct {
	VertexA [3]string
	VertexB [3]string
	VertexC [3]string
}

func getDXF(UnitList []*AngleUnit, fileName string) error {
	f, err := os.OpenFile("./save/"+fileName[:len(fileName)-4]+".dxf", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	content := make([]byte, 0, 5*1024*1024)
	content = append(content, []byte("0\nSECTION\n2\nENTITIES\n")...)

	for _, face := range UnitList {
		content = append(content, []byte("0\n3DFACE\n8\n0\n")...)
		content = append(content, []byte(fmt.Sprintf("10\n%s\n20\n%s\n30\n%s\n", face.VertexA[0], face.VertexA[1], face.VertexA[2]))...)
		content = append(content, []byte(fmt.Sprintf("11\n%s\n21\n%s\n31\n%s\n", face.VertexB[0], face.VertexB[1], face.VertexB[2]))...)
		content = append(content, []byte(fmt.Sprintf("12\n%s\n22\n%s\n32\n%s\n", face.VertexC[0], face.VertexC[1], face.VertexC[2]))...)
		content = append(content, []byte(fmt.Sprintf("13\n%s\n23\n%s\n33\n%s\n", face.VertexC[0], face.VertexC[1], face.VertexC[2]))...)
	}
	content = append(content, []byte("0\nENDSEC\n0\nEOF\n")...)
	f.Write(content)
	return nil
}

func getNormal(face *AngleUnit) (x, y, z float32) {
	temp := make([]float32, 0, 9)
	for _, val := range face.VertexA {
		temp = append(temp, getFloat(val))
	}
	for _, val := range face.VertexB {
		temp = append(temp, getFloat(val))
	}
	for _, val := range face.VertexC {
		temp = append(temp, getFloat(val))
	}

	x = (temp[4]-temp[1])*(temp[8]-temp[2]) - (temp[5]-temp[2])*(temp[7]-temp[1])
	y = (temp[6]-temp[0])*(temp[5]-temp[2]) - (temp[3]-temp[0])*(temp[8]-temp[2])
	z = (temp[3]-temp[0])*(temp[7]-temp[1]) - (temp[6]-temp[0])*(temp[4]-temp[1])
	return
}

func getFloat(s string) float32 {
	result, err := strconv.ParseFloat(s, 4)
	if err != nil {
		return 0
	}
	return float32(result)
}

func getSTL(UnitList []*AngleUnit, fileName string) error {
	content := make([]byte, 0, 5*1024*1024)
	content = append(content, []byte("solid filenamestl "+fileName[:len(fileName)-4]+".stl\n")...)

	for _, face := range UnitList {
		x, y, z := getNormal(face)
		content = append(content, []byte(fmt.Sprintf("facet normal %f %f %f\n", x, y, z))...)
		content = append(content, []byte("outer loop\n")...)

		content = append(content, []byte(fmt.Sprintf("vertex %s %s %s\n", face.VertexA[0], face.VertexA[1], face.VertexA[2]))...)
		content = append(content, []byte(fmt.Sprintf("vertex %s %s %s\n", face.VertexB[0], face.VertexB[1], face.VertexB[2]))...)
		content = append(content, []byte(fmt.Sprintf("vertex %s %s %s\n", face.VertexC[0], face.VertexC[1], face.VertexC[2]))...)

		content = append(content, []byte("endloop\nendfacet\n")...)
	}
	content = append(content, []byte("endsolid filenamestl "+fileName[:len(fileName)-4]+".stl\n")...)

	f, err := os.OpenFile("./save/"+fileName[:len(fileName)-4]+".stl", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	f.Write(content)
	return err
}

func GetFile(redisDB *redis.Client, fileName string) error {
	result, err := redisDB.LRange("./temp/"+fileName, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	redisDB.Del("./temp/" + fileName)

	UnitList := make([]*AngleUnit, 0, len(result)/9)

	for i := 0; i < len(result); i += 9 {
		unit := AngleUnit{}
		unit.VertexA[0] = result[i]
		unit.VertexA[1] = result[i+1]
		unit.VertexA[2] = result[i+2]

		unit.VertexB[0] = result[i+3]
		unit.VertexB[1] = result[i+4]
		unit.VertexB[2] = result[i+5]

		unit.VertexC[0] = result[i+6]
		unit.VertexC[1] = result[i+7]
		unit.VertexC[2] = result[i+8]

		UnitList = append(UnitList, &unit)
	}

	if !strings.HasSuffix(fileName, ".dxf") {
		err = getDXF(UnitList, fileName)
	} else {
		err = getSTL(UnitList, fileName)
	}

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
