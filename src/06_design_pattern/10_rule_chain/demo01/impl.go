// @Time : 2023/8/29 19:04
// @Author : xiaoweiwei
// @File : impl

package demo01

import "context"

type baseRuleChain struct {
	next RuleChain
}

func (b *baseRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	panic(any("not implement"))
}

func (b *baseRuleChain) Next() RuleChain {
	return b.next
}

func (b *baseRuleChain) applyNext(ctx context.Context, params map[string]interface{}) error {
	if b.Next() != nil {
		return b.Next().Apply(ctx, params)
	}
	return nil
}
