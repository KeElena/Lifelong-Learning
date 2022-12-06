package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//加载待发送请求的URI
	URI, _ := url.ParseRequestURI("http://127.0.0.1:8081/usr")
	//构建文本io
	io := strings.NewReader(`{"id":1120,"name":"keqing"}`)

	resPost, _ := http.Post(URI.String(), "json", io)

	data2, _ := ioutil.ReadAll(resPost.Body)
	fmt.Println(string(data2))

}