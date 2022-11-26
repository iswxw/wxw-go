/*
@Time : 2022/5/12 09:52
@Author : weixiaowei
@File : demo02_geek
*/
package main

import "fmt"

// https://time.geekbang.org/column/article/386238
// 做菜流程 说明 模板方法模式

type Cooker interface {
	fire()    // 开火
	cooke()   // 做菜
	outFire() // 关火
}

// 定义一个抽象类——————————————————————————————————————
type CookMenu struct {
}

func (cm CookMenu) fire() {
	fmt.Println("开火")
}

func (cm CookMenu) outFire() {
	fmt.Println("关火")
}

// 做菜，交给具体的子类实现
func (cm CookMenu) cooke() {}

// 做菜执行者
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outFire()
}

// 定义子类1 实现具体做的菜品——————————————————————————————————

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

// 定义子类2 实现具体做的菜品——————————————————————————————————
type ChaoJiDan struct {
	CookMenu
}

func (*ChaoJiDan) cooke() {
	fmt.Println("炒鸡蛋")
}
