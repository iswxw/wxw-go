// @Time : 2023/8/29 19:35
// @Author : xiaoweiwei
// @File : check_token

package demo01

import (
	"context"
	"fmt"
)

type CheckTokenRule struct {
	baseRuleChain
}

func NewCheckTokenRule(next RuleChain) RuleChain {
	return &CheckTokenRule{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckTokenRule) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验 token 是否合法
	token, _ := params["token"].(string)
	if token != "myToken" {
		return fmt.Errorf("invalid token: %s", token)
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check token rule err post process...")
		return err
	}

	fmt.Println("check token rule common post process...")
	return nil
}
