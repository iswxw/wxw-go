/*
@Time : 2022/3/29 11:10
@Author : weixiaowei
@File : demo03_student_test struct 反序列化耗时测试
@相关文章：https://geektutu.com/post/hpg-sync-pool.html
*/
package demo03_student_benchmark_test

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte `json:"remark"`
}

// 定义一个池对象
var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

// 基准测试反序列化
// BenchmarkUnmarshal-8   	   14155	     79393 ns/op
func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

// 基准测试对象池反序列化
// BenchmarkUnmarshalWithPool-8   	   14403	     83573 ns/op
func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
