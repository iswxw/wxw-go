/*
@Time: 2022/7/23 15:34
@Author: wxw
@File: BM6_HasCycle
*/
package _1_list

import (
	dto2 "src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func TestBM6(t *testing.T) {

}

// hasCycle 判断链表是否有环：快慢指针法
func hasCycle(head *dto2.ListNode) bool {
	// write code here

	// 边界处理
	if head == nil || head.Next == nil {
		return false
	}

	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
