// @Time : 2023/8/29 19:03
// @Author : xiaoweiwei
// @File : interface

package demo01

import "context"

// 定义接口
type RuleChain interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	Next() RuleChain
}
