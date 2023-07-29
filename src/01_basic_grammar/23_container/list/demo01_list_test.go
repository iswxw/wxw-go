// @Time : 2023/5/30 17:29
// @Author : xiaoweiwei
// @File : demo01_list_test

package list

import (
	"container/list"
	"fmt"
	"testing"
)

// TestHelloWorld 01_list 是一个链表结构
func TestListHelloWorld(t *testing.T) {
	// Create a new 01_list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)

	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through 01_list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
