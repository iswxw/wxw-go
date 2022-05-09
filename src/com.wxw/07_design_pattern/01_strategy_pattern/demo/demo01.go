/*
@Time : 2022/1/19 22:44
@Author : weixiaowei
@File : demo 文档：http://books.studygolang.com/go-patterns/behavioral/strategy.html
*/
package main

// 策略
type Operator interface {
	Apply(int, int) int
}

// 执行策略的类：具体策略的执行者
type Operation struct {
	Operator Operator
}

// 调用策略中的方法
func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

// 执行
func main() {
	add := Operation{
		Operator: Addition{},
	}
	add.Operate(3, 5) // 8
}

// 策略一 加法
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

// 策略二 乘法
type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
