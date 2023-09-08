/*
@Time: 2023/9/8 23:04
@Author: wxw
@File: BM15_deleteDuplicates_test
*/
package _1_list

import (
	"src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func TestBM15(t *testing.T) {

}

// 重复的元素保留一个
// 题目描述：删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
// 思路：双指针
func deleteDuplicates(head *dto.ListNode) *dto.ListNode {
	// write code here
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
