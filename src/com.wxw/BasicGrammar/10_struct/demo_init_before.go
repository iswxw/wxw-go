/*
@Time: 2021/10/6 15:27
@Author: wxw
@File: demo_init_before
*/
package main

import "fmt"

type person1 struct {
	name string
	age  int64
	sex  bool
}

func main() {
	var p1 person1
	fmt.Printf("p1 = %#v\n", p1)
	// p1 = main.person1{name:"", age:0, sex:false}
}
