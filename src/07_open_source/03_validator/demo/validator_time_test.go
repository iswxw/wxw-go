// @Time : 2023/5/22 20:50
// @Author : xiaoweiwei
// @File : validator_time

package demo

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

type User struct {
	Name     string `validate:"min=6,max=10"`
	Age      int    `validate:"min=1,max=100"`
	Birthday string `validate:"required,datetime=2006-01-02 15:04:05"`
	Content  string `validate:"required"`
	Phone    string `validate:"omitempty"`
}

// 出自：https://blog.51cto.com/u_15301988/5133385
// omitempty: https://www.cnblogs.com/MyUniverse/p/15227003.html
func TestCheckTime(t *testing.T) {
	validate := validator.New()

	u1 := User{
		Name:     "wxw-go",
		Age:      18,
		Birthday: "2022-03-29 22:22:10",
	}
	err := validate.Struct(u1)
	fmt.Println("u1 = ", err)

	u2 := User{
		Name:     "wxw-go",
		Age:      18,
		Birthday: "2022-03-29",
	}
	err1 := validate.Struct(u2)
	fmt.Println("u2 = ", err1)
}
