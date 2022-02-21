/*
@Time : 2022/2/14 18:45
@Author : weixiaowei
@File : solution_time
*/
package main

import (
	"encoding/json"
	"log"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	ct := time.Time(t)
	if y := ct.Year(); y < 2016 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return json.Marshal("")
	}
	b = ct.AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	t2 := time.Time(t)
	return t2.Format(timeFormart)
}

type Person struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Birthday   *Time  `json:"birthday,omitempty"`
	UpdateTime *Time  `json:"updateTime,omitempty"`
}

// [golang 自定义time.Time json输出格式] https://www.cnblogs.com/xiaofengshuyu/p/5664654.html
func main() {
	now := Time(time.Now())
	log.Println("current = ", now)

	//src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51","updateTime":""}`
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	p := &Person{}
	if err := json.Unmarshal([]byte(src), p); err != nil {
		log.Println("err:", err)
	}

	js, _ := json.Marshal(p)

	log.Println(string(js))
}
