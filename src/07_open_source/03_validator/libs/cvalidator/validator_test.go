// @Time : 2023/3/24 15:54
// @Author : xiaoweiwei
// @File : validator_test

package cvalidator

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"testing"
)

type DayQuitOrder struct {
	OrderOrgan string `validate:"required,order_organ"  json:"order_organ" desc:"出单机构"`
	PolicyNo   string `validate:"required,policy_no,upper_number"  json:"policy_no" desc:"保单号"`
}

func TestCheckStruct(t *testing.T) {

	ctx := context.Background()
	v := validator.New()
	vv := GetCV()
	err := vv.DecorateValidator(v, CustomTags)
	if err != nil {
		println("err = ", err)
		return
	}

	dayQuitOrder := &DayQuitOrder{
		OrderOrgan: "测试机构",
		PolicyNo:   "ATAYmMCE2123B00166AA",
	}

	err = v.StructCtx(ctx, dayQuitOrder)
	println("err = ", err.Error())
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			println("err = ", err)
			return
		}
		tErrs := errs.Translate(vv.Trans)

		errSlice := make([]string, 0)
		desc := GetQuitOrderFieldDesc()

		for f, errInfo := range tErrs {
			str := strings.Split(f, ".")
			if strings.HasSuffix(errInfo, "为必填字段") {
				errSlice = append(errSlice, fmt.Sprintf("%s字段为空", desc[str[1]]))
			} else {
				errSlice = append(errSlice, fmt.Sprintf("%s字段有误", desc[str[1]]))
			}
		}

		println(fmt.Sprintf("errSlice = %v", errSlice))
	}

}

func GetQuitOrderFieldDesc() map[string]string {
	m := GetStructField2Tag(DayQuitOrder{})
	r := make(map[string]string)
	for k, v := range m {
		r[k] = v.Get("desc")
	}
	return r
}

func GetStructField2Tag(structName interface{}) map[string]reflect.StructTag {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	fieldNum := t.NumField()
	result := make(map[string]reflect.StructTag)
	for i := 0; i < fieldNum; i++ {
		result[t.Field(i).Name] = t.Field(i).Tag
	}
	return result
}
