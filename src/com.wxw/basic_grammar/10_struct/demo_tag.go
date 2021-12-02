/*
@Time: 2021/10/7 0:43
@Author: wxw
@File: demo_tag
*/
package main

import (
	"encoding/json"
	"fmt"
)

// Teacher 学生
type Teacher struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问

}

func main() {
	s1 := Teacher{
		ID:     1,
		Gender: "男",
		name:   "半颗糖",
	}
	// 序列化
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
