/*
@Time: 2021/10/6 23:26
@Author: wxw
@File: demo_method
*/
package _struct

import (
	"fmt"
	"testing"
)

// Person 结构体
type Person struct {
	name string
	age  int64
}

// NewPerson 构造函数
func NewPerson(name string, age int64) *Person {
	return &Person{name: name, age: age}
}

// Dream Person 做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s 的梦想是学好Go语言！\n", p.name)
}

// SetAge 指针类型得方法
func (p *Person) SetAge(newAge int64) {
	p.age = newAge
}

// SetAge1 值类型的方法
func (p Person) SetAge1(newAge int64) {
	fmt.Printf("p1 = %#v, newAge = %v \n", p, newAge)
	p.age = newAge
}

func TestMethod(t *testing.T) {
	p1 := NewPerson("半颗糖", 18)
	p1.Dream()

	p1.SetAge(19)
	fmt.Printf("p1 = %#v \n", p1)

	p1.SetAge1(20)
	fmt.Printf("p1 = %#v \n", p1)

}
