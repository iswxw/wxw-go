/*
@Time: 2023/9/8 23:53
@Author: wxw
@File: BM16_deleteDuplicates_test
*/
package _1_list

import (
	"src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func Test16(t *testing.T) {

}

// 重复的元素全部删除
// 题目描述：
// 给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
// 给出的链表为1→1→1→2→3 返回2→3.

// 方法一：递归
func deleteDuplicate(head *dto.ListNode) *dto.ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicate(head.Next)
		return head
	} else {
		cur := head.Next
		for cur != nil && head.Val == cur.Val {
			cur = cur.Next
		}
		return deleteDuplicate(cur)
	}
}

// 方法二：遍历
func deleteDuplicate01(head *dto.ListNode) *dto.ListNode {
	if head == nil {
		return head
	}

	dummy := &dto.ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
