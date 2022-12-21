// @Time : 2022/12/22 00:18
// @Author : xiaoweiwei
// @File : dog

package dog

import (
	"log"
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/constant"
)

type Dog struct {
}

func (d *Dog) Name() string {
	return "dog"
}

func (d *Dog) Live() {
	d.Eat()
	d.Drink()
	d.Shit()
	d.Pee()
	d.Run()
	d.Howl()
}

func (d *Dog) Eat() {
	log.Println(d.Name(), "eat")
}

func (d *Dog) Drink() {
	log.Println(d.Name(), "drink")
}

func (d *Dog) Shit() {
	log.Println(d.Name(), "shit")
}

func (d *Dog) Pee() {
	log.Println(d.Name(), "pee")
}

func (d *Dog) Run() {
	log.Println(d.Name(), "run")
	_ = make([]byte, 16*constant.Mi)
}

func (d *Dog) Howl() {
	log.Println(d.Name(), "howl")
}
