/*
@Time: 2022/6/3 16:46
@Author: wxw
@File: BM1_ReverseList
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代实现
func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	var newHead *ListNode
	for pHead != nil {
		pNext := pHead.Next  // 保留未反转的链表
		pHead.Next = newHead // 节点反转
		newHead = pHead      // 更新已反转的节点
		pHead = pNext        // 更新当前节点
	}
	return newHead
}

// 递归实现
func ReverseList01(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := ReverseList01(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
