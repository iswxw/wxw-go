/*
@Time : 2022/2/14 18:45
@Author : weixiaowei
@File : solution_time
*/
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyTime time.Time

func (m *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	return json.Unmarshal(data, (*time.Time)(m))
}

func main() {
	s := `[
	{"name":"test1","expireAt":"2050-12-31T00:00:00Z"},
	{"name":"test2","expireAt":""}
	]`
	var result []struct {
		Name     string
		ExpireAt MyTime
	}
	if err := json.Unmarshal([]byte(s), &result); err != nil {
		fmt.Println(err)
	}
	for _, v := range result {
		fmt.Println(time.Time(v.ExpireAt))
	}
}
