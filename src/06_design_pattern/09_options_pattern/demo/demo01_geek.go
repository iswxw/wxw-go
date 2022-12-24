/*
@Time : 2022/5/12 11:09
@Author : weixiaowei
@File : 09_options_pattern
*/
package main

import "time"

// 定义一个连接器
type Collection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	caching bool
	timeout time.Duration
}

// 定义选择模式
type Option interface {
	apply(options *options)
}

// 定义一个选择函数
type optionFunc func(*options)

// 应用函数的实现
func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(
		func(o *options) {
			o.timeout = t
		})
}

func WithCaching(cache bool) Option {
	return optionFunc(
		func(o *options) {
			o.caching = cache
		})
}

// 创建一个连接器
func NewCollect(addr string, opts ...Option) (*Collection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, opt := range opts {
		opt.apply(&options)
	}

	return &Collection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}
