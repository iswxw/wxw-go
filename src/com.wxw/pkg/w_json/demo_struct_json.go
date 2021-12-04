/*
@Time: 2021/11/27 21:49
@Author: wxw
@File: demo_struct_json
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// User 定义一个结构体
type User struct {
	UserName string    `json:"user_name,omitempty"` // omitempty 表示忽略空值
	Age      int       `json:"age"`
	Age1     int       `json:"-"` //`json:"-"` 表示不进行序列化,忽略这个字段
	Gender   string    `json:"gender,string"`
	Birthday time.Time `json:"birthday" `
}

func NewUser() *User {
	return &User{
		UserName: "Java半颗糖",
		Age:      18,
		Gender:   "男",
	}
}

func NewEmptyUser() *User {
	return &User{}
}

func main() {
	user := NewUser()
	fmt.Printf("user: %#v\n", user)

	// marshal
	byteUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal("json marshal error:", err)
	}
	fmt.Printf(" user: %#v\n", byteUser)

	// unmarshal
	user1 := NewEmptyUser()
	if err = json.Unmarshal(byteUser, &user1); err != nil {
		log.Fatal("json unmarshal error:", err)
	}
	fmt.Printf(" user: %#v\n", user1)
}