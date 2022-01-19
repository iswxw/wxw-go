/*
@Time : 2022/1/19 22:44
@Author : weixiaowei
@File : demo01 文档：http://books.studygolang.com/go-patterns/behavioral/strategy.html
*/
package main

// 控制器
type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

func main() {
	add := Operation{
		Operator: Addition{},
	}
	add.Operate(3, 5) // 8
}

// 策略一
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

// 策略二
type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
