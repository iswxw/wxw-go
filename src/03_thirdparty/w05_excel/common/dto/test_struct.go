/*
@Time : 2022/1/26 00:14
@Author : weixiaowei
@File : User
*/
package dto

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Price     string  `json:"price"`
	UnitPrice float64 `json:"unitPrice"`
}
