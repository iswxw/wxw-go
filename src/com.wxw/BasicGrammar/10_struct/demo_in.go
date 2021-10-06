/*
@Time: 2021/10/6 14:45
@Author: wxw
@File: demo_success
*/
package main

import "fmt"

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}

// 嵌入结构体
func main() {
	user1 := User{
		Name:   "半颗糖",
		Gender: "男",
		Address: Address{
			Province: "北京",
			City:     "北京",
		},
	}
	fmt.Printf("user1=%#v\n", user1)
	// user1=main.User{Name:"半颗糖", Gender:"男", Address:main.Address{Province:"北京", City:"北京"}}

}
