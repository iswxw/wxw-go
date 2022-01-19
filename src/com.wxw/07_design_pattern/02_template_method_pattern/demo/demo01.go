/*
@Time : 2022/1/19 23:10
@Author : weixiaowei
@File : demo01 模板方法模式
*/
package main

import "fmt"

type Template interface {
	fun1()
	fun2()
	Result()
}

// 抽象结构体
type Funcs struct {
	temp Template
}

// 抽象结构体部分实现接口
func (r *Funcs) Result() {
	r.temp.fun2()
	r.temp.fun1()
}

//A具体实现。继承自抽象结构体的方法+自身实现的方法=实现Template接口
type ConcreteA struct {
	//会继承抽象结构体中的方法
	Funcs
}

func (c *ConcreteA) fun1() {
	fmt.Println("A类实现fun1")
}

func (c *ConcreteA) fun2() {
	fmt.Println("A类实现fun2")
}

//B具体实现。继承自抽象结构体的方法+自身实现的方法=实现Template接口
type ConcreteB struct {
	//会继承抽象结构体中的方法
	Funcs
}

func (c *ConcreteB) fun1() {
	fmt.Println("B类实现fun1")
}

func (c *ConcreteB) fun2() {
	fmt.Println("B类实现fun2")
}

func main() {
	instance := ConcreteB{}
	fmt.Printf("%P \n", instance.Funcs)
	ta := Funcs{temp: &instance}
	ta.Result()
}
