/*
@Time: 2021/9/21 17:59
@Author: wxw
@File: demo 结构体
@Link: https://www.runoob.com/go/go-structures.html
*/
package main

import "fmt"

// person 定义一个结构体
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "张三"
	p1.city = "北京"
	p1.age = 12
	fmt.Printf("p1=%v\n", p1)  //p1={张三 北京 12}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"张三", city:"北京", age:12}

	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}
