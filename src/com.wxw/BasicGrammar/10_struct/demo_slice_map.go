/*
@Time: 2021/10/7 0:56
@Author: wxw
@File: demo_slice_map
*/
package main

import "fmt"

type Person1 struct {
	name   string
	age    int8
	dreams []string
}

// SetDreams 通过指针定义方法
func (p *Person1) SetDreams(dreams []string) {
	p.dreams = dreams
}

// SetDreams1 正确的做法是在方法中使用传入的slice的拷贝进行结构体赋值。
func (p1 *Person1) SetDreams1(dreams1 []string) {
	p1.dreams = make([]string, len(dreams1))
	copy(p1.dreams, dreams1)
}
func main() {
	p1 := Person1{name: "半颗糖", age: 18}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)

	// 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams) // ?

	p1.SetDreams1(data)
	data[1] = "睡觉1"
	fmt.Println(p1.dreams) //

}
