/*
@Time: 2023/8/20 17:38
@Author: wxw
@File: BM8_find_k_tail 返回该链表中倒数第k个节点
*/
package _1_list

import (
	dto2 "src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func TestBM8(t *testing.T) {

}

// FindKthToTail 返回链表中倒数第k个节点：双指针法
// 大致思路：第一个指针先移动k步，然后第二个指针再从头开始，这个时候这两个指针同时移动，当第一个指针到链表的末尾的时候，返回第二个指针即可
func FindKthToTail(pHead *dto2.ListNode, k int) *dto2.ListNode {
	// write code here

	p1, p2 := pHead, pHead

	// 第一个指针先移动k步
	for k > 0 && p2 != nil {
		k--
		p2 = p2.Next
	}

	// 如果 k 大于链表长度,直接返回，此时:k>len(p2)
	if k > 0 {
		return nil
	}

	// p2 还没走完，则符合快慢指针的条件
	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
