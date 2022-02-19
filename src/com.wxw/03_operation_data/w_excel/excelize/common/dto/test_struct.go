/*
@Time : 2022/1/26 00:14
@Author : weixiaowei
@File : User
*/
package dto

import "time"

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Price     string  `json:"price"`
	UnitPrice float64 `json:"unitPrice"`
}

type ExcelTest struct {
	Name     string    `json:"name" name:"名称" index:"0"`
	Age      int64     `json:"age" name:"年龄" index:"1"`
	Score    int64     `json:"score" name:"分数" index:"2"`
	Birthday time.Time `json:"birthday" name:"生日" index:"4"`
}
