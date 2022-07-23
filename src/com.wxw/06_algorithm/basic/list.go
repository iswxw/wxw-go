/*
@Time: 2022/7/23 14:18
@Author: wxw
@File: list
*/
package main

import "fmt"

// 定义一个链表结构
type ListNode struct {
	val  int
	next *ListNode
}

// 头插法
func InitList(head *ListNode, num int) *ListNode {
	var tail *ListNode
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < num; i++ {
		var node = ListNode{val: i}
		node.next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
	}
	return tail
}

func Show(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next //移动指针
	}
}

func main() {
	node := &ListNode{val: 0}
	Show(InitList(node, 10))
}
