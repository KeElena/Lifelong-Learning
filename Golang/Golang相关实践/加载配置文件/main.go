package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type config struct {
	Mysql `ini:"mysql"`
}

type Mysql struct {
	Port     string `ini:"port"`
	Host     string `ini:"host"`
	User     string `ini:"user"`
	Password string `ini:"password"`
}

func loadIni(addr string, data interface{}) error {
	configType := reflect.TypeOf(data)

	if configType.Kind() != reflect.Ptr {
		err := errors.New("not pointer")
		return err
	}
	//pointer -> val
	if configType.Elem().Kind() != reflect.Struct {
		err := errors.New("not struct")
		return err
	}
	ini, errR := ioutil.ReadFile(addr)
	if errR != nil {
		return errR
	}
	lineSlice := strings.Split(string(ini), "\n")

	var structName string
	for i, line := range lineSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "[") {
			var sectionName string
			if strings.HasSuffix(line, "]") {
				sectionName = strings.TrimSpace(line[1 : len(line)-1])
				if len(sectionName) == 0 {
					return fmt.Errorf("line:%d error", i+1)
				} else {
					for i := 0; i < configType.Elem().NumField(); i++ {
						configStruct := configType.Elem().Field(i)
						if sectionName == configStruct.Tag.Get("ini") {
							structName = configStruct.Name
							break
						}
					}
				}
			} else {
				return fmt.Errorf("line:%d error", i+1)
			}
		} else {
			val := strings.Split(line, "=")
			if strings.Index(line, "=") == -1 {
				return fmt.Errorf("line %d error", i+1)
			}
			if val[0] == "" {
				return fmt.Errorf("line %d error", i+1)
			}

			configValue := reflect.ValueOf(data)
			structValue := configValue.Elem().FieldByName(structName)

			var fieldName string
			for i := 0; i < structValue.NumField(); i++ {
				field := structValue.Type().Field(i)
				if field.Tag.Get("ini") == val[0] {
					fieldName = field.Name
					break
				}
			}

			fieldVal := structValue.FieldByName(fieldName)
			switch fieldVal.Kind() {
			case reflect.String:
				fieldVal.SetString(strings.TrimSpace(val[1]))
			}
		}
	}
	return nil
}

func main() {
	var a config
	fmt.Println(loadIni("/home/keqing/桌面/go/src/load_ini/mysql.ini", &a))
	fmt.Println(a)
}
