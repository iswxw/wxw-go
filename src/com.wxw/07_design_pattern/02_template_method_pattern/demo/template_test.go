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
