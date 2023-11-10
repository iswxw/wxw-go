// @Time : 2022/12/22 00:10
// @Author : xiaoweiwei
// @File : muridae

package muridae

import (
	"src/com.wxw/project_actual/src/02_advance_grammar/w_performance_tuning/pprof/animal"
)

// Muridae 鼠科动物
type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
