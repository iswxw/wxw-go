/*
@Time: 2021/10/7 0:13
@Author: wxw
@File: demo_success
*/
// 结构体的继承
package main

import "fmt"

// Animal 结构体
type Animal struct {
	name string
}

// 指针类型的方法
func (a *Animal) move() {
	fmt.Printf("%s 会动！\n", a.name)
}

// Dog 结构体
type Dog struct {
	Feet    int8
	*Animal // 通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s 会汪汪叫 \n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
