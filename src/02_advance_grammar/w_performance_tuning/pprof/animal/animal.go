// @Time : 2022/12/22 00:08
// @Author : xiaoweiwei
// @File : animal

package animal

import (
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal/canidae/dog"
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal/canidae/wolf"
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal/felidae/cat"
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal/felidae/tiger"
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},
		&cat.Cat{},
		&tiger.Tiger{},
		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
