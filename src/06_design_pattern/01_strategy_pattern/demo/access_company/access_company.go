// @Time : 2023/4/3 16:17
// @Author : xiaoweiwei
// @File : consts

package access_company

import (
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/access_company/base"
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/access_company/impl"
)

const (
	NumberPICC = 1001
	NumberCIPC = 1002
)

var MapAccessCompany = map[int]base.IAccessCompany{
	NumberPICC: impl.NewAccessPICC(),
	NumberCIPC: impl.NewAccessCIPC(),
}
