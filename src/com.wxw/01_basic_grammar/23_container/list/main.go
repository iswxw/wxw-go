package main

import (
	"container/list"
	"fmt"
)

// list 是一个 双向链表
func main() {
	l := list.New()
	fmt.Println("新的List:",l)
	fmt.Println("List长度:",l.Len())
}

