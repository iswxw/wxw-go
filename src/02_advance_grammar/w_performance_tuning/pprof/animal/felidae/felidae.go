// @Time : 2022/12/22 00:13
// @Author : xiaoweiwei
// @File : felidea

package felidae

import "src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal"

// Felidae 猫科动物
type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
