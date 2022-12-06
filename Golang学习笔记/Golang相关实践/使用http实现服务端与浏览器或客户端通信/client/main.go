package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//resp, err := http.Get("http://127.0.0.1:8080/test?name=who&age=18")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//s, _ := ioutil.ReadAll(resp.Body)
	//resp.Close()
	//fmt.Println(string(s))
	URI, _ := url.ParseRequestURI("http://127.0.0.1:8080/test")

	data := url.Values{}
	data.Set("name", "我")
	data.Set("age", "18")

	URI.RawQuery = data.Encode()
	fmt.Println(URI)

	fileObj, _ := os.Open("/home/keqing/桌面/go/src/http/server/hello.html")
	reader := bufio.NewReader(fileObj)

	req, err := http.NewRequest("GET", URI.String(), reader)

	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	Client := http.Client{Transport: tr}

	resp, err := Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	s, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))
}
