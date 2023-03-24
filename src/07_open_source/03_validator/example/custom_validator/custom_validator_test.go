/*
@Time: 2023/3/15 22:04
@Author: wxw
@File: example
*/
package custom_validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

// 出自：https://blog.51cto.com/u_15301988/5133385
func TestValidator(t *testing.T) {

	validate := validator.New()
	// 注册自定义标签
	_ = validate.RegisterValidation("required_if", CheckRequiredIf)

	f1 := RegisterForm{
		Name: "djd",
		Age:  18,
	}
	err := validate.Struct(f1)
	if err != nil {
		fmt.Println("f1 = ", err)
	}

	f2 := RegisterForm{
		Name: "dj",
		Age:  18,
	}
	err = validate.Struct(f2)
	if err != nil {
		fmt.Println("f2 = ", err)
	}

}
