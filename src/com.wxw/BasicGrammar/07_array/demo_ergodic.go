/*
@Time: 2021/9/21 11:23
@Author: wxw
@File: demo_02 数组的遍历
*/
package main

import "fmt"

func main() {

	var a = [...]string{"北京", "上海", "深圳"}

	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	// 方法2：for range 遍历
	for index,value:= range a{
		fmt.Println(index,value)
	}
}