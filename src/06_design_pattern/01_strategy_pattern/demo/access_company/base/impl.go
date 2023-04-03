// @Time : 2023/4/3 16:07
// @Author : xiaoweiwei
// @File : access_company

package base

import (
	"context"
)

type AccessCompany struct {
	impl IAccessCompany
}

func (a *AccessCompany) SetTest(ctx context.Context) {
	a.impl.SetTest(ctx)
}
