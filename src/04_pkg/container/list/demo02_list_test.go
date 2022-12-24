package main

import (
	"container/list"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4) // 插入链表最后
	fmt.Println("4插入链表最后", e4.Value)

	e1 := l.PushFront(1) // 插入链表最前
	fmt.Println("1插入链表最前", e1.Value)

	e3 := l.InsertBefore(3, e4) // 插入指定元素 e4 前面
	fmt.Println("3插入指定元素 e4 前面", e3.Value)

	e2 := l.InsertAfter(2, e1) // 插入指定元素 e1 之后
	fmt.Println("2插入指定元素 e1 之后", e2.Value)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d", e.Value)
	}
	fmt.Println()

	// Output:
	// 1
	// 2
	// 3
	// 4
}
