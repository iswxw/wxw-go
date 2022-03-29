/*
@Time : 2022/1/18 09:27
@Author : weixiaowei
@File : demo
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	go func() {
		// 循环打印
		for {
			log.Println(Add("Hello world"))
		}
	}()

	if err := http.ListenAndServe("0.0.0.0:6060", nil); err != nil {
		log.Println("err:", err)
		return
	}
}

var data []string

func Add(str string) string {
	data1 := []byte(str)
	sData := string(data1)
	data = append(data, sData)
	return sData
}
