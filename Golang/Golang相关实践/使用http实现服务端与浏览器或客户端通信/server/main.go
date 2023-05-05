package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hellohandlerFunc(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		w.Write([]byte("not content"))
		return
	}
	w.Write(content)
}

func testHandlerFunc(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	fmt.Println(val.Get("age"))
	fmt.Println(val.Get("name"))
	fmt.Println(r.Method)
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(data))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello", hellohandlerFunc)
	http.HandleFunc("/test", testHandlerFunc)
	http.ListenAndServe("127.0.0.1:8080", nil)

}
