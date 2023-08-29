// @Time : 2023/8/29 19:36
// @Author : xiaoweiwei
// @File : check_authorized_status

package demo01

import (
	"context"
	"errors"
	"fmt"
)

type CheckAuthorizedStatus struct {
	baseRuleChain
}

func NewCheckAuthorizedStatus(next RuleChain) RuleChain {
	return &CheckAuthorizedStatus{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAuthorizedStatus) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验是否已认证
	if authorized, _ := params["authorized"].(bool); !authorized {
		return errors.New("not authorized yet")
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check authorized status rule err post process...")
		return err
	}

	fmt.Println("check authorized statuse rule common post process...")
	return nil
}
