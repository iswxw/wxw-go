/*
@Time: 2021/11/28 15:47
@Author: wxw
@File: demo_1
*/
package main

import "fmt"

const (
	UnPay = iota
	HadPay
	Delivery
	Finish
)

var orderState = map[int]string{
	UnPay:    "未支付",
	HadPay:   "已支付",
	Delivery: "配送中",
	Finish:   "已完成",
}

// 场景一：根据不同的错误码显示对应错误消息，比如 200 -> 正常。
// 场景二：根据不同状态显示对应的文案。这个场景很常见，比如数据库保存状态，用的 tinyint 类型，显示给用户的是文本，所以需要进行转换。
func main() {
	fmt.Println(OrderStateMap(0))
	fmt.Println(OrderStateSwitch(0))
}

// OrderStateMap map 实现
func OrderStateMap(state int) string {
	return orderState[state]
}

// OrderStateSwitch switch 实现
func OrderStateSwitch(state int) string {
	var stateDesc = ""

	switch state {
	case UnPay:
		stateDesc = "未支付"
	case HadPay:
		stateDesc = "已支付"
	case Delivery:
		stateDesc = "配送中"
	case Finish:
		stateDesc = "已完成"
	}

	return stateDesc
}
