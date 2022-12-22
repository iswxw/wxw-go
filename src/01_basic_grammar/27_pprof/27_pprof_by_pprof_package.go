/*
@Time: 2022/12/22 10:10
@Author: wxw
@File: 27_pprof
*/
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

var datas []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}

// 通过net/http/pprof包 生成性能分析文件
// localhost:6060
func main() {
	go func() {
		for {
			log.Println(Add("https://github.com/iswxw"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
