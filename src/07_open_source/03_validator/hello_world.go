// @Time : 2023/3/15 17:38
// @Author : xiaoweiwei
// @File : hello_world

package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `validate:"min=6,max=10"`
	Age  int    `validate:"min=1,max=100"`
}

// 出自：https://blog.51cto.com/u_15301988/5133385
func main() {
	validate := validator.New()

	u1 := User{
		Name: "wxw-go",
		Age:  18,
	}
	err := validate.Struct(u1)
	fmt.Println("u1 = ", err)

	u2 := User{
		Name: "wxw",
		Age:  101,
	}
	err = validate.Struct(u2)
	fmt.Println("u2 = ", err)
}
