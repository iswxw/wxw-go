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

	user.Birthday = time.Now()
	// marshal
	byteUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal("json marshal error:", err)
	}
	fmt.Printf(" user: %s\n", string(byteUser))

	// unmarshal
	// byteUser1 := []byte(`{"user_name":"Java半颗糖","age":18,"gender":"\"男\"","birthday":"0001-01-01T00:00:00Z"}`)
	byteUserData := []byte(`{"user_name":"Java半颗糖","age":18,"gender":"\"男\"","birthday":"2021-04-16T00:00:01Z"}`)
	user1 := NewEmptyUser()
	if err = json.Unmarshal(byteUserData, &user1); err != nil {
		log.Fatal("json unmarshal error:", err)
	}

	fmt.Printf(" user: %s\n", user1)
}
