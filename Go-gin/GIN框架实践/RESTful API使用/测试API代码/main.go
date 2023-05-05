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
	URI, _ := url.ParseRequestURI("http://127.0.0.1:8081/hello")
	//构建文本io
	io := strings.NewReader("")
	//发送Get请求
	resGet, _ := http.Get(URI.String())
	//发送Post请求
	resPost, _ := http.Post(URI.String(), "", nil)
	//构建Put请求
	reqPut, _ := http.NewRequest(http.MethodPut, URI.String(), io)
	//构建Del请求
	reqDel, _ := http.NewRequest(http.MethodDelete, URI.String(), io)
	//发送Put请求
	resPut, _ := http.DefaultClient.Do(reqPut)
	//发送Del请求
	resDel, _ := http.DefaultClient.Do(reqDel)

	//取值
	data1, _ := ioutil.ReadAll(resGet.Body)
	fmt.Println(string(data1))
	data2, _ := ioutil.ReadAll(resPost.Body)
	fmt.Println(string(data2))
	data3, _ := ioutil.ReadAll(resPut.Body)
	fmt.Println(string(data3))
	data4, _ := ioutil.ReadAll(resDel.Body)
	fmt.Println(string(data4))
}
