/*
@Time: 2022/9/5 22:32
@Author: wxw
@File: tencent_test
*/
package demo01_reflect

import "testing"

func BenchmarkDeleteSliceElms(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	elms := []interface{}{uint64(1), uint64(3), uint64(5), uint64(7), uint64(9)}
	for i := 0; i < b.N; i++ {
		_ = DeleteSliceElms(slice, elms...)
	}
}

func BenchmarkDeleteU64liceElms(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	elms := []uint64{1, 3, 5, 7, 9}
	for i := 0; i < b.N; i++ {
		_ = DeleteU64liceElms(slice, elms...)
	}
}

func BenchmarkNtohl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ntohl([]byte{0x7f, 0, 0, 0x1})
	}
}

func BenchmarkNtohlNotUseBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NtohlNotUseBinary([]byte{0x7f, 0, 0, 0x1})
	}
}
