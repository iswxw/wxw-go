package main

import (
	"container/heap"
	"fmt"
	"reflect"
	"testing"
)

// 自定义实现一个堆
type myHeap []int

// 实现 Less
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// 实现 Swap
func (h *myHeap) Swap(i, j int) {
	(*h)[j], (*h)[i] = (*h)[i], (*h)[j]
	return
}

// 实现 Len
func (h *myHeap) Len() int {
	return len(*h)
}

// 实现 Pop
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[:h.Len()-1]
	return
}

// 实现 Push
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

//========================开始单元测试================================

// 默认每次从堆顶取出最小元素
func Test(t *testing.T) {
	h := new(myHeap)
	for i := 20; i > 0; i-- {
		h.Push(i)
	}

	heap.Init(h)
	for h.Len() > 0 {
		pop := heap.Pop(h)
		fmt.Println(fmt.Sprintf("%s : %v;", reflect.TypeOf(pop), pop))
	}

}
