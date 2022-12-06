package ParaseObj

import (
	"bufio"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

type Vertex struct {
	VertexA string
	VertexB string
	VertexC string
}

func ParseObj(reader *bufio.Reader, redisDB *redis.Client, fileName string) error {
	units := make([]*Vertex, 0, 1024)
	faces := make([]int, 0, 1024*5)
	for {
		temp, _, err := reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				err = nil
				break
			} else {
				return err
			}
		}

		if strings.HasPrefix(string(temp), "v ") {
			unit := &Vertex{}
			cache := strings.Fields(string(temp)[2:])
			if len(cache) > 3 {
				return fmt.Errorf("format mismatch")
			}
			unit.VertexA = cache[0]
			unit.VertexB = cache[1]
			unit.VertexC = cache[2]
			fmt.Println(string(temp))
			units = append(units, unit)
		}
		if strings.HasPrefix(string(temp), "f ") {
			cache := strings.Fields(string(temp)[2:])
			if len(cache) > 3 {
				return fmt.Errorf("format mismatch")
				return err
			}
			for _, str := range cache {
				temp := strings.Split(str, "/")
				v, err := strconv.Atoi(temp[0])
				if err != nil {
					return err
				}
				faces = append(faces, v)
			}
		}
	}

	pipe := redisDB.Pipeline()
	defer pipe.Close()

	for _, idx := range faces {
		err := pipe.LPush(fileName, units[idx-1].VertexA, units[idx-1].VertexB, units[idx-1].VertexC).Err()
		if err != nil {
			return err
		}
	}
	pipe.Expire(fileName, time.Second)
	_, err := pipe.Exec()
	if err != nil {
		return err
	}

	return nil
}
