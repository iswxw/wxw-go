// @Time : 2023/8/29 18:57
// @Author : xiaoweiwei
// @File : demo01_rule_chain_test

package demo01

import (
	"context"
	"fmt"
	"testing"
)

// 主题：责任链模式
// 相关资料：
func TestDemo01(t *testing.T) {
	checkAuthorizedRule := NewCheckAuthorizedStatus(nil)
	checkAgeRule := NewCheckAgeRule(checkAuthorizedRule)
	checkTokenRule := NewCheckTokenRule(checkAgeRule)

	if err := checkTokenRule.Apply(context.Background(), map[string]interface{}{
		"token": "myToken",
		"age":   1,
	}); err != nil {
		// 校验未通过，终止发奖流程
		t.Error(err)
		return
	}

	// 通过前置校验流程，进行奖励发放
	//sendReward(ctx,params)
	fmt.Println("通过前置校验流程，进行奖励发放")
}
