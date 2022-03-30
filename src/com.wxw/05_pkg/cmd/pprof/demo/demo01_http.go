/*
@Time : 2022/1/18 09:27
@Author : weixiaowei
@File : demo
*/
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

var datas []string

// 相关文章：https://www.jianshu.com/p/f4690622930d
func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
	go func() {
		for {
			log.Println(Add("Hello world"))
		}
	}()
	// 访问地址：http://localhost:6060/debug/pprof/
	http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}
