// @Time : 2022/12/17 16:54
// @Author : xiaoweiwei
// @File : options

package worker_pool

type Option func(*Pool)

func WithBlock(block bool) Option {
	return func(p *Pool) {
		p.block = block
	}
}

func WithPreAllocWorkers(preAlloc bool) Option {
	return func(p *Pool) {
		p.preAlloc = preAlloc
	}
}
