/*
@Time: 2023/8/20 17:53
@Author: wxw
@File: BM9_removeNthFromEnd_test 给定一个链表，删除链表的倒数第 n 个节点并返回链表的头指针
*/
package _1_list

import (
	dto2 "src/com.wxw/project_actual/src/05_algorithm/common/dto"
)

// removeNthFromEnd 双指针法
func removeNthFromEnd(head *dto2.ListNode, n int) *dto2.ListNode {
	// write code here

	// 标记一下头节点
	dummy := &dto2.ListNode{}
	dummy.Next = head

	// 定义两个指针
	p1, p2 := dummy, head

	// 快指针先走n步
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next

	return dummy.Next
}
