/*
@Time: 2023/3/15 22:44
@Author: wxw
@File: custom_param_error
*/
package custom_error

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// processErr 错误处理
func processParamErr(err error) {
	if err == nil {
		return
	}

	// 参数错误
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		fmt.Println("param error:", invalid)
		return
	}
}
