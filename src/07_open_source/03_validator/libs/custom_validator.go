// @Time : 2023/3/15 19:39
// @Author : xiaoweiwei
// @File : custom_validator

package main

import "github.com/go-playground/validator/v10"

//validate.RegisterValidation("custom tag name", customFunc)
//// NOTES: using the same tag name as an existing function
////        will overwrite the existing one

// Structure
func customFunc(fl validator.FieldLevel) bool {

	if fl.Field().String() == "invalid" {
		return false
	}

	return true
}

func main() {

}
