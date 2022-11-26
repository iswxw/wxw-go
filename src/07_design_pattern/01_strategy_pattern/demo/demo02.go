/*
@Time : 2022/1/19 22:54
@Author : weixiaowei
@File : demo02
*/
package main

import "fmt"

// 策略一
type Car struct {
}

func (r *Car) Go() {
	fmt.Println("use car")
}

// 策略二
type Bicycle struct {
}

func (r *Bicycle) Go() {
	fmt.Println("use Bicycle")
}

// 管控策略
type Vehicle interface {
	Go()
}
type Traveler struct {
	impl Vehicle
}

// 设置策略
func (r *Traveler) SetVehicle(i Vehicle) {
	r.impl = i
}

// 执行策略
func (r *Traveler) Go() {
	r.impl.Go()
}

func main() {
	traveler := Traveler{}
	traveler.SetVehicle(&Car{})
	traveler.Go()

	traveler.SetVehicle(&Bicycle{})
	traveler.Go()
}
