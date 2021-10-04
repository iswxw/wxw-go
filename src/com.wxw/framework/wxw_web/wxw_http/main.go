/*
@Time: 2021/9/21 22:47
@Author: wxw
@File: main
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// http://127.0.0.1:9000/hello
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("http serve failed, err: %v \n", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 给文件中写入内容
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}
