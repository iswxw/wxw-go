/*
@Time : 2022/2/26 17:53
@Author : weixiaowei
@File : demo01 https://www.jianshu.com/p/dbd7a5bc21e9
*/
package main

import "fmt"

// 状态模式1
// 当今社会，论坛贴吧很多，我们也会加入感兴趣的论坛，偶尔进行发言，但有时却会发现不能发帖了，原来是昨天的某个帖子引发了口水战，被举报了。
// 这里就用论坛发帖为例，简单用代码描述一下：https://www.jianshu.com/p/dbd7a5bc21e9

// 在状态模式中将对象在每一个状态下的行为和状态转移语句封装在一个个状态类中，通过这些状态类来分散冗长的条件转移语句，让系统具有更好的灵活性和可扩展性。
// 假设有三种状态，normal(正常），restricted(受限)，closed(封号)，判断依据是一个健康值

func main() {
	account := NewAccount(-11)
	account.View()
	account.Post()
	account.Comment()
}

// 定义一个用户
type Account struct {
	State       ActionState
	HealthValue int
}

// 用户实例
func NewAccount(health int) *Account {
	a := &Account{
		HealthValue: health,
	}
	a.changeState()
	return a
}

func (a *Account) View() {
	a.State.View()
}

func (a *Account) Comment() {
	a.State.Comment()
}
func (a *Account) Post() {
	a.State.Post()
}

// 状态管理器
type ActionState interface {
	View()
	Comment()
	Post()
}

// 关闭
type CloseState struct {
}

func (c *CloseState) View() {
	fmt.Println("账号被封，无法看帖")
}

func (c *CloseState) Comment() {
	fmt.Println("抱歉，你的健康值小于-10，不能评论")
}
func (c *CloseState) Post() {
	fmt.Println("抱歉，你的健康值小于0，不能发帖")
}

// 禁止
type RestrictedState struct {
}

func (r *RestrictedState) View() {
	fmt.Println("正常看帖")
}

func (r *RestrictedState) Comment() {
	fmt.Println("正常评论")
}
func (r *RestrictedState) Post() {
	fmt.Println("抱歉，你的健康值小于0，不能发帖")
}

// 正常
type NormalState struct {
}

func (n *NormalState) View() {
	fmt.Println("正常看帖")
}

func (n *NormalState) Comment() {
	fmt.Println("正常评论")
}
func (n *NormalState) Post() {
	fmt.Println("正常发帖")
}

func (a *Account) changeState() {
	if a.HealthValue <= -10 {
		a.State = &CloseState{}
	} else if a.HealthValue > -10 && a.HealthValue <= 0 {
		a.State = &RestrictedState{}
	} else if a.HealthValue > 0 {
		a.State = &NormalState{}
	}
}

///给账户设定健康值
func (a *Account) SetHealth(value int) {
	a.HealthValue = value
	a.changeState()
}
