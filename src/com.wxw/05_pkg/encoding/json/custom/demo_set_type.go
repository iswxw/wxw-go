/*
@Time: 2022/2/22 8:42
@Author: wxw
@File: demo_set_type
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

// 定制JSON序列化 https://www.jianshu.com/p/6015912a120d
func main() {

	movie := Movie{
		ID:       1,
		CreateAt: time.Now(),
		Title:    "Casablanca",
		Year:     int32(time.Now().Year()),
		Runtime:  Runtime(30),
		Genres:   []string{"DD"},
		Version:  200,
	}

	marshal, err := json.Marshal(movie)
	if err != nil {
		log.Println("err = ", err)
	}
	log.Println("result = ", string(marshal))
}

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`
	Title    string    `json:"title"`
	Year     int32     `json:"year,omitempty"`
	//使用Runtime类型取代int32，注意omitempty还是能生效的
	Runtime Runtime  `json:"runtime,omitempty,string"`
	Genres  []string `json:"genres,omitempty"`
	Version int32    `json:"version"`
}

//申明Runtime类型，其底层是int32类型（和movie中的字段一样）
type Runtime int32

//实现MarshalJSON()方法，这样就实现了json.Marshaler接口。
func (r Runtime) MarshalJSON() ([]byte, error) {
	//生成一个字符串包含电影时长
	jsonValue := fmt.Sprintf("%d mins", r)

	//使用strconv.Quote()函数封装双引号。为了在JSON中以字符串对象输出，需要用双引号。
	quotedJSONValue := strconv.Quote(jsonValue)
	//将字符串转为[]byte返回
	return []byte(quotedJSONValue), nil
}
