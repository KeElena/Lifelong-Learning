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
	} 

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}