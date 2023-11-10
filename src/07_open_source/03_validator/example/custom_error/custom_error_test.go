/*
@Time: 2023/3/15 22:25
@Author: wxw
@File: custom_error
*/
package custom_error

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"testing"
)

// 参数错误校验
func TestErrorValidator(t *testing.T) {
	validate := validator.New()

	fmt.Println("参数错误")
	err := validate.Struct(1)
	processErr(err)

	fmt.Println("\n校验错误")
	err = validate.VarWithValue("name", "age", "eqcsfield") //  将两个变量进行对比
	processErr(err)
}

// 中文错误处理
func TestChineseError(t *testing.T) {
	users := &Users{
		Name:   "admin",
		Age:    12,
		Passwd: "123",
		Code:   "123456",
	}

	// 中文翻译器
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	// 校验器
	validate := validator.New()

	// 注册翻译器到校验器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err = ", err)
	}
	err = validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("err = ", err.Translate(trans))
			return
		}
	}
}

// 校验错误带tag提示
func TestTagTips(t *testing.T) {
	// 参数校验
	validate := validator.New()

	userBasic := &UserBasic{
		Name: "wxw-go",
		Age:  18,
	}

	// 自定义校验
	//validate.RegisterValidation("required_if", CheckRequiredIf)
	if err := validate.Struct(userBasic); err != nil {
		// 返回我们tag中的提示语
		ProcessErrAddTag(userBasic, err)
		return
	}
}

type Users struct {
	Name   string `form:"name" json:"name" validate:"required"`
	Age    uint8  `form:"age" json:"age" validate:"required,gt=18"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

type UserBasic struct {
	UserId     uint64 `gorm:"column:user_id" json:"userId"`
	UserNumber string `gorm:"column:user_number" json:"userNumber"`
	Name       string `validate:"required" tag:"姓名不能为空"`
	Age        uint8  `validate:"lt=0|gt=150" tag:"年龄不合法"`
	PassWord   string `gorm:"column:password" json:"password"`
	PhoneNum   string `validate:"required_if" tag:"手机号格式不正确"`
	Email      string `validate:"email" tag:"email为空或格式不正确"`
}
