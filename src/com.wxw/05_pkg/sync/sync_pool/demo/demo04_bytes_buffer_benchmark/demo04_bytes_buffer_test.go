/*
@Time : 2022/3/29 11:36
@Author : weixiaowei
@File : demo04_bytes_buffer_benchmark
https://geektutu.com/post/hpg-sync-pool.html
*/
package demo04_bytes_buffer_benchmark

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

// 这个例子创建了一个 bytes.Buffer 对象池，而且每次只执行一个简单的 Write 操作，存粹的内存搬运工，耗时几乎可以忽略。
// 而内存分配和回收的耗时占比较多，因此对程序整体的性能影响更大。

// 执行指令：go test -bench=. -benchmem
var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}
