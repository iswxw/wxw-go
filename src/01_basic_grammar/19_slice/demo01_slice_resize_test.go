// @Time : 2023/5/30 16:46
// @Author : xiaoweiwei
// @File : demo01_slice_resize_test

package main

import (
	"fmt"
	"testing"
)

// 资料：https://www.yuque.com/fcant/go/mkb552#LgESr
// 当v<=1.17时： 当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍； 原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。
// 当v= 1.19时： 当前 cap > 2倍的原cap时，新容量 = 当前cap；
//   - 当cap < 256 时新 slice 容量变成原来的 2 倍；
//   - 当cap >= 256 容量为512时触发扩容时,新 slice 容量变成原来的 1.25 倍；
func TestResize(t *testing.T) {
	s := make([]int, 0)
	oldCap := cap(s)
	for i := 1; i <= 2048; i++ {
		s = append(s, i)
		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("[before append %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", i-1, oldCap, len(s), newCap)
			oldCap = newCap
		}
	}
}

// 源码
// 源代码地址：go 1.19.5 src/runtime/slice.go:82

// growslice handles slice growth during append.
// It is passed the slice element type, the old slice, and the desired new minimum capacity,
// and it returns a new slice with at least that capacity, with the old data
// copied into it.
// The new slice's length is set to the old slice's length,
// NOT to the new requested capacity.
// This is for codegen convenience. The old slice's length is used immediately
// to calculate where to write new values during an append.
// TODO: When the old backend is gone, reconsider this decision.
// The SSA backend might prefer the new length or to return only ptr/cap and save stack space.
//func growslice(et *_type, old slice, cap int) slice {

//	newcap := old.cap
//	doublecap := newcap + newcap
//	if cap > doublecap {
//		newcap = cap
//	} else {
//		const threshold = 256
//		if old.cap < threshold {
//			newcap = doublecap
//		} else {
//			// Check 0 < newcap to detect overflow
//			// and prevent an infinite loop.
//			for 0 < newcap && newcap < cap {
//				// Transition from growing 2x for small slices
//				// to growing 1.25x for large slices. This formula
//				// gives a smooth-ish transition between the two.
//				newcap += (newcap + 3*threshold) / 4
//			}
//			// Set newcap to the requested cap when
//			// the newcap calculation overflowed.
//			if newcap <= 0 {
//				newcap = cap
//			}
//		}
//	}
//
//	return slice{p, old.len, newcap}
//}
