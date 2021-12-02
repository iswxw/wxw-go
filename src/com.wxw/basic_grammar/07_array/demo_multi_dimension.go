/*
@Time: 2021/9/21 17:09
@Author: wxw
@File: demo_multi_dimension 二维数组
*/
package main

import "fmt"

// 多维数组只有第一层可以使用...来让编译器推导数组长度
// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆

	// 二维数组的遍历
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

}
