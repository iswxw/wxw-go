package cvalidator

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"strings"
)

var CustomTags = []CustomTag{
	{
		// float不超过两位小数校验
		Tag: "float_two",
		Func: func(ctx context.Context, fl validator.FieldLevel) bool {
			f := fl.Field().String()
			ff, err := strconv.ParseFloat(f, 64)
			if err != nil {
				return false
			}
			str := strings.Split(fmt.Sprintf("%f", ff), ".")
			if len(str) < 2 {
				return false
			}
			num, err := strconv.Atoi(str[1])
			if err != nil {
				return false
			}
			if num%10000 > 0 {
				return false
			}
			return true
		},
		Translation: "{0}不能超过两位小数",
		Override:    false,
	},
	{ // 大写字母+数字校验
		Tag: "upper_number",
		Func: func(ctx context.Context, fl validator.FieldLevel) bool {
			f := fl.Field().String()
			result, err := regexp.MatchString(UpperNumber, f)
			if err != nil {
				return false
			}
			return result
		},
		Translation: "{0}不符合都是大写字母和数字的规则",
		Override:    false,
	},
	{
		// 出单机构校验
		Tag: "order_organ",
		Func: func(ctx context.Context, fl validator.FieldLevel) bool {
			f := fl.Field().String()
			if strings.HasSuffix(f, "公司") || strings.HasSuffix(f, "部") {
				return true
			}
			return false
		},
		Translation: "{0}不符合出单机构名称规则",
		Override:    false,
	},
	{
		// 保单号校验，末尾不是000000结束
		Tag: "policy_no",
		Func: func(ctx context.Context, fl validator.FieldLevel) bool {
			f := fl.Field().String()
			if strings.HasSuffix(f, "000000") {
				return false
			}
			return true
		},
		Translation: "{0}不符合保单号结尾不能为000000的规则",
		Override:    false,
	},
}
