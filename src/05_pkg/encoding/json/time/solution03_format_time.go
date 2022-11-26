/*
@Time : 2022/2/21 22:11
@Author : weixiaowei
@File : solution03_format_time
*/
package main

import (
	"encoding/json"
	"log"
	util2 "src/com.wxw/project_actual/src/05_pkg/util"
	"time"
)

type Student struct {
	Id         int64             `json:"id"`
	Name       string            `json:"name"`
	Birthday   util2.FormatTime  `json:"birthday"`
	UpdateTime *util2.FormatTime `json:"updateTime,omitempty"`
}

// 此时 omitempty 生效必须使用指针类型
// [golang 自定义time.Time json输出格式] https://www.cnblogs.com/xiaofengshuyu/p/5664654.html
func main() {
	now := util2.FormatTime(time.Now())
	log.Println("current = ", now)

	//src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51","updateTime":""}`
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	p := &Student{}
	if err := json.Unmarshal([]byte(src), p); err != nil {
		log.Println("err:", err)
	}

	js, _ := json.Marshal(p)
	log.Println(string(js))
}
