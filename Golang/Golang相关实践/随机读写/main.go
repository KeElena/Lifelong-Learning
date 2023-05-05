package main

import (
	"fmt"
	"os"
)

func randRead(fileAddr string) {

	fileRead, errO := os.Open(fileAddr)
	if errO != nil {
		fmt.Println(errO)
	}
	defer fileRead.Close()

	var data = make([]byte, 30)
	fileRead.Seek(2, 0)
	for {
		n, _ := fileRead.Read(data)
		fmt.Print(string(data[:n]))
		if n < len(data) {
			break
		}
	}
}

func randWrite(fileAddr string) {

	fileObj, errO := os.OpenFile(fileAddr, os.O_CREATE|os.O_RDWR, 0777)
	if errO != nil {
		fmt.Println(errO)
	}
	defer fileObj.Close()

	fileObj.Seek(3, 0)
	data := make([]byte, 20)
	s := make([]byte, 10)
	for {
		n, _ := fileObj.Read(s)
		data = append(data, s[:n]...)
		if n < len(s) {
			break
		}
	}

	str := "\ngolang"
	fileObj.Seek(3, 0)
	fileObj.WriteString(str + string(data))
}

func main() {
	//randRead("/home/keqing/桌面/go/src/hello world/demo.txt")
	//randWrite("/home/keqing/桌面/go/src/hello world/demo.txt")
}
