/*
@Time: 2022/3/18 16:01
@Author: wxw
@File: demo01_no_state.go
*/
package main

import "fmt"

// 存在的问题：
//  1. 看帖和发帖方法中都包含状态判断语句，以判断在该状态下是否具有该方法以及在特定状态下该方法如何实现，导致代码非常冗长，可维护性较差；
//  2. 系统扩展性较差，如果需要增加一种新的状态，如hot状态（活跃用户，该状态用户发帖积分增加更多），需要对原有代码进行大量修改，扩展起来非常麻烦。
func main() {
	account01 := NewAccount01(0)
	account01.Comment()
}

type AccountState int

const (
	NORMAL     AccountState = iota //正常0
	RESTRICTED                     //受限
	CLOSED                         //封号
)

type Account01 struct {
	State       AccountState
	HealthValue int
}

func NewAccount01(health int) *Account01 {
	a := &Account01{
		HealthValue: health,
	}
	a.changeState()
	return a
}

///看帖
func (a *Account01) View() {
	if a.State == NORMAL || a.State == RESTRICTED {
		fmt.Println("正常看帖")
	} else if a.State == CLOSED {
		fmt.Println("账号被封，无法看帖")
	}

}

///评论
func (a *Account01) Comment() {
	if a.State == NORMAL || a.State == RESTRICTED {
		fmt.Println("正常评论")
	} else if a.State == CLOSED {
		fmt.Println("抱歉，你的健康值小于-10，不能评论")
	}

}

///发帖
func (a *Account01) Post() {
	if a.State == NORMAL {
		fmt.Println("正常发帖")
	} else if a.State == RESTRICTED || a.State == CLOSED {
		fmt.Println("抱歉，你的健康值小于0，不能发帖")
	}
}

func (a *Account01) changeState() {
	if a.HealthValue <= -10 {
		a.State = CLOSED
	} else if a.HealthValue > -10 && a.HealthValue <= 0 {
		a.State = RESTRICTED
	} else if a.HealthValue > 0 {
		a.State = NORMAL
	}
}

///给账户设定健康值
func (a *Account01) SetHealth(value int) {
	a.HealthValue = value
	a.changeState()
}
