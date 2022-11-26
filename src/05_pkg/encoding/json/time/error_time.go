/*
@Time : 2022/2/14 18:44
@Author : weixiaowei
@File : error_time
*/
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	s := `[
	{"name":"test1","expireAt":"2050-12-31T00:00:00Z"},
	{"name":"test2","expireAt":""}
	]`
	var result []struct {
		Name     string
		ExpireAt time.Time
	}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range result {
		fmt.Println(v.ExpireAt)
	}
}
