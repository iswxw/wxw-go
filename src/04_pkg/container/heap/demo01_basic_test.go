package main

import (
	"container/heap"
	"fmt"
	"testing"
)

type Person struct {
	Name  string  //名字
	Age   int     //年龄
	Money float64 //身价
}

type heapTool []*Person

// Less、Swap、Len 实现 Sort 的功能
// 比较
func (h *heapTool) Less(i, j int) bool {
	return (*h)[i].Age < (*h)[j].Age
}

// 交换
func (h *heapTool) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 长度
func (h *heapTool) Len() int {
	return len(*h)
}

// 出堆
func (h *heapTool) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 入堆
func (h *heapTool) Push(v interface{}) {
	*h = append(*h, v.(*Person))
}

func TestPopAndPush(t *testing.T) {

	personList := new(heapTool)
	personList.Push(&Person{
		Name:  "小明",
		Age:   20,
		Money: 100000.99,
	})
	personList.Push(&Person{
		Name:  "小施",
		Age:   30,
		Money: 1002341.99,
	})

	personList.Push(&Person{
		Name:  "小康",
		Age:   10,
		Money: 200.99,
	})

	personList.Push(&Person{
		Name:  "老施",
		Age:   50,
		Money: 10343240000.99,
	})

	personList.Push(&Person{
		Name:  "老康",
		Age:   70,
		Money: 10340.99,
	})
	personList.Push(&Person{
		Name:  "老明",
		Age:   80,
		Money: 13240000.99,
	})
	personList.Push(&Person{
		Name:  "老林",
		Age:   90,
		Money: 10340000.99,
	})

	// 默认小顶堆
	heap.Init(personList)
	for personList.Len() > 0 {
		pop := heap.Pop(personList)
		fmt.Println(fmt.Sprintf("%v,%v岁，资产：%v", pop.(*Person).Name, pop.(*Person).Age, pop.(*Person).Money))
	}

}
