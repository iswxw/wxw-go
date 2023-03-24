/*
@Time: 2023/3/15 22:03
@Author: wxw
@File: custom_validator
*/
package custom_validator

import (
	"github.com/go-playground/validator/v10"
)

type RegisterForm struct {
	Name string `validate:"required_if"`
	Age  int    `validate:"min=18"`
}

// CheckRequiredIf 判断字符串是否是对称的
func CheckRequiredIf(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == reverseString(value)
}

// reverseString 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
