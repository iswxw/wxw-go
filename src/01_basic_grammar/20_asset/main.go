/*
@Time : 2022/4/29 10:56
@Author : weixiaowei
@File : main
*/
package main

import "fmt"

func main() {
	var container = []string{"zero", "one", "two"}
	value, ok := interface{}(container).([]string)
	if ok {
		fmt.Println("value1 = ", value)
	}
	fmt.Println("value2 = ", value)
}
