/*
@Time: 2021/9/21 17:59
@Author: wxw
@File: case 结构体
@Link: https://www.runoob.com/go/go-structures.html
*/
package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

// TestEmptyStruct 空结构体，详见：https://www.yuque.com/fcant/go/gq99os
func TestEmptyStruct(t *testing.T) {
	var a int
	var b string
	var c bool
	var d [3]int32
	var e []string
	var f map[string]bool
	var s struct{}

	fmt.Println(
		// 可以查看一个类型的实例所占用的存储空间的字节数
		unsafe.Sizeof(a),
		unsafe.Sizeof(b),
		unsafe.Sizeof(c),
		unsafe.Sizeof(d),
		unsafe.Sizeof(e),
		unsafe.Sizeof(f),
		unsafe.Sizeof(s),
	)
}

// TestInitBefore 初始化之前
func TestInitBefore(t *testing.T) {
	var p person
	fmt.Printf("p = %#v\n", p)
	// p1 = main.person1{name:"", age:0, sex:false}
}

// TestInitAfter 初始化之后
func TestInitAfter(t *testing.T) {
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

// TestInStruct 嵌入结构体
func TestInStruct(t *testing.T) {
	user1 := User{
		Name:   "半颗糖",
		Gender: "男",
		Address: Address{
			Province: "北京",
			City:     "北京",
		},
	}
	fmt.Printf("user1=%#v\n", user1)
	// user1=main.User{Name:"半颗糖", Gender:"男", Address:main.Address{Province:"北京", City:"北京"}}

}

// person 定义一个结构体
type person struct {
	name string
	city string
	age  int8
}

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address Address
}
