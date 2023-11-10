/*
@Time: 2023/3/15 22:25
@Author: wxw
@File: custom_error
*/
package custom_error

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// processErr 错误处理
func processErr(err error) {
	if err == nil {
		return
	}

	// 参数错误
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return
	}

	// 校验错误
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationError
	for _, validationErr := range validationErrs {
		fmt.Println("validator error:", validationErr)
	}
}

// processErrAddTag 参数校验不合规则使用tag提示
func ProcessErrAddTag(u interface{}, err error) bool {
	if err == nil {
		return true
	}

	// 参数错误
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return false
	}

	// 校验错误
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationError
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() //获取是哪个字段不符合格式
		typeOf := reflect.TypeOf(u)

		// 如果是指针，获取其属性
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		field, ok := typeOf.FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("tag")        // 获取field对应的tag值
			fmt.Println(fieldName + ":" + errorInfo) // 返回错误
			return false
		} else {
			fmt.Println("缺失tag")
			return false
		}
	}
	return true
}

// CheckRequiredIf 判断值是否是对的
func CheckRequiredIf(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	fmt.Println("________________")
	fmt.Println("fl.Field = ", fl.Field())
	fmt.Println("fl.FieldName = ", fl.FieldName())
	fmt.Println("fl.GetTag = ", fl.GetTag())
	fmt.Println("________________")
	return value == ""
}
