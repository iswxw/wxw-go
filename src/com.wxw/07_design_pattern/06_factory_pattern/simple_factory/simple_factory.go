/*
@Time: 2022/4/9 21:15
@Author: wxw
@File: factory
*/
package main

// 解析规则配置的接口
type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

type yamlRuleConfigParser struct{}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me!")
}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me!")
}

// 实现规则配置
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	default:
		return nil

	}
}
