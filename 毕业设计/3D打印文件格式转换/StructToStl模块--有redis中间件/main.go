package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var redisDB *redis.Client

type AngleUnit struct {
	VertexA [3]string
	VertexB [3]string
	VertexC [3]string
}

func initRedis() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0})
	_, err := redisDB.Ping().Result()
	if err != nil {
		return
	}
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

func GetFile(redisDB *redis.Client, fileName string) error {
	result, err := redisDB.LRange("./temp/"+fileName, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}

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

	err = getSTL(UnitList, fileName)
	if err != nil {
		return err
	}

	return nil
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

	f, err := os.OpenFile("./demo.stl", os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	f.Write(content)
	return err
}

func main() {
	initRedis()
	GetFile(redisDB, "芭芭拉.stl")
}
