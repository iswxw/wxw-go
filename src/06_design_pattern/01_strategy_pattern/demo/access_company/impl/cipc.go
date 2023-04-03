// @Time : 2023/4/3 16:08
// @Author : xiaoweiwei
// @File : cipc

package impl

import (
	"context"
	"src/com.wxw/project_actual/src/06_design_pattern/01_strategy_pattern/demo/access_company/base"
)

type AccessCIPC struct {
	base.IAccessCompany
}

func NewAccessCIPC() *AccessCIPC {
	return &AccessCIPC{}
}

func (a *AccessCIPC) SetTest(ctx context.Context) {

}
