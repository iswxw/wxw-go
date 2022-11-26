/*
@Time: 2021/10/10 11:18
@Author: wxw
@File: demo_lwz https://www.liwenzhou.com/posts/Go/12_interface/
*/
package main

import "fmt"

func main() {
	// 值引用
	var move Mover
	var jack = dog{}
	move = jack

	var link = dog{}
	move = link
	move.move()

	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}

type Mover interface {
	move()
}

type dog struct {
}

// 实现
func (d dog) move() {
	fmt.Println("狗会动")
}
