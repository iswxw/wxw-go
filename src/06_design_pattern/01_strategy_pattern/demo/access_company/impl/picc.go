// @Time : 2023/4/3 16:08
// @Author : xiaoweiwei
// @File : picc

package impl

import (
	"context"
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/access_company/base"
)

type AccessPICC struct {
	base.IAccessCompany
}

func NewAccessPICC() *AccessPICC {
	return &AccessPICC{}
}

func (a *AccessPICC) SetTest(ctx context.Context) {
	println("my name is AccessPICC")
}
