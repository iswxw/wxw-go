// @Time : 2022/12/22 00:16
// @Author : xiaoweiwei
// @File : canidae

package canidae

import "src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal"

// Canidae 犬科动物
type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
