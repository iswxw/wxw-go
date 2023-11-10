/*
@Time : 2021/11/23 00:20
@Author : wxw
@File : for_update_value
*/
package _for

import "fmt"

/**
  Go 通过for循环遍历数组并修改值
    - 通过range获取数组的值 -> 不能修改原数组中结构体的值：
    - 通过range获取数组下标 -> 可以修改原数组中结构体的值：
  总结：想要通过遍历数组方式去改变其值时，需要通过数组下标去操作（示例2）
  link: https://blog.csdn.net/qq_37102984/article/details/117850578
*/

func main() {
	type User struct {
		Name string
	}

	userArr := []User{
		{Name: "ZhangSan"},
		{Name: "LiSi"},
	}

	// 示例1：通过range获取数组的值 -> 不能修改原数组中结构体的值：
	// 输出：userName1:{ZhangSan} userName1:{LiSi}；原数组值并未改变！
	for _, value := range userArr {
		value.Name = "WangWu"
	}

	for _, value := range userArr {
		fmt.Println("userName1: ", value)
	}

	// 示例2：通过range获取数组下标 -> 可以修改原数组中结构体的值：
	// 输出：userName2:{WangWu} userName2:{WangWu}；原数组值被成功改变！
	for k, _ := range userArr {
		userArr[k].Name = "WangWu"
	}

	for _, userName2 := range userArr {
		fmt.Println("userName2: ", userName2)
	}

}
