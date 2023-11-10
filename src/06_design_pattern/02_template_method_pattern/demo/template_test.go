/*
@Time : 2022/5/12 10:04
@Author : weixiaowei
@File : template_test
*/
package main

import (
	"fmt"
	"testing"
)

// 实现方式一
func TestTemplate_demo01(t *testing.T) {
	// 具体实现1
	A := &ConcreteA{}
	ta := Funcs{temp: A}
	ta.Result()

	fmt.Println("\n=====> 执行另外一个具体实现")
	// 具体实现2
	B := &ConcreteB{}
	tb := Funcs{temp: B}
	tb.Result()

}

// 实现方式二
func TestTemplate_demo02(t *testing.T) {
	// 做西红柿
	xihongshi := &XiHongShi{}
	doCook(xihongshi)

	fmt.Println("\n=====> 做另外一道菜")
	// 做炒鸡蛋
	chaojidan := &ChaoJiDan{}
	doCook(chaojidan)

}

// ==============================具体实现============================

type Template interface {
	// 具体子类实现
	fun1()
	fun2()

	// 抽象类实现
	Result()
}

// 抽象结构体
type Funcs struct {
	temp Template
}

// 抽象结构体部分实现
func (r *Funcs) Result() {
	r.temp.fun2()
	r.temp.fun1()
}

// A具体实现。继承自抽象结构体的方法+自身实现的方法=实现Template接口
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

// B具体实现。继承自抽象结构体的方法+自身实现的方法=实现Template接口
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
