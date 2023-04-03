// @Time : 2023/4/3 16:08
// @Author : xiaoweiwei
// @File : access_company_test

package access_company

import (
	"context"
	"testing"
)

func TestName(t *testing.T) {
	company := MapAccessCompany[NumberPICC]
	ctx := context.Background()
	company.SetTest(ctx)
}
