/*
@Time: 2021/12/4 23:43
@Author: wxw
@File: benchmark
*/
package main

import "fmt"

// 基准测试 源代码

// 测试方法 case_1
func HandleType(flag int) string {
	fmt.Printf("HandleType: %v", flag)
	switch flag {
	case 0:
		return "Add"
	case 1:
		return "Sub"
	case 2:
		return "Multiply"
	case 3:
		return "Division"
	}
	return "NotExist"
}
