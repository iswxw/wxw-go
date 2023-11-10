// @Time : 2023/8/29 19:36
// @Author : xiaoweiwei
// @File : check_age

package demo01

import (
	"context"
	"fmt"
)

type CheckAgeRule struct {
	baseRuleChain
}

func NewCheckAgeRule(next RuleChain) RuleChain {
	return &CheckAgeRule{
		baseRuleChain: baseRuleChain{
			next: next,
		},
	}
}

func (c *CheckAgeRule) Apply(ctx context.Context, params map[string]interface{}) error {
	// 校验 age 是否合法
	age, _ := params["age"].(int)
	if age < 18 {
		return fmt.Errorf("invalid age: %d", age)
	}

	if err := c.applyNext(ctx, params); err != nil {
		// err post process
		fmt.Println("check age rule err post process...")
		return err
	}

	fmt.Println("check age rule common post process...")
	return nil
}
