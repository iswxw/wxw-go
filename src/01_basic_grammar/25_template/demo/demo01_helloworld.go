// @Time : 2022/8/22 15:24
// @Author : xiaoweiwei
// @File : demo01_helloworld

package main

import (
	"html/template"
	"log"
	"net/http"
)

// 案例blogs：https://www.cnblogs.com/zhangyafei/p/12482244.html

func sayHello(w http.ResponseWriter, r *http.Request) {

	basePath := "src/com.wxw/01_basic_grammar/25_template/tmpl/demo01_helloworld.tmpl"

	// 解析模板
	t, err := template.ParseFiles(basePath)
	if err != nil {
		log.Printf("Parse template failed, err%v\n", err)
		return
	}
	// 渲染模板
	name := "技术能量站"
	err = t.Execute(w, name)
	if err != nil {
		log.Printf("render template failed, err%v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("http server start failed,err:%v\n", err)
	}
}
