/*
@Time: 2023/8/20 16:47
@Author: wxw
@File: BM4_merge_order_list
@link:https://www.nowcoder.com/practice/d8b6b4358f774294a89de2a6ac4d9337
*/
package _1_list

import (
	dto2 "src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func TestBM4(t *testing.T) {

}

// Merge 合并两个有序链表
func Merge(pHead1 *dto2.ListNode, pHead2 *dto2.ListNode) *dto2.ListNode {
	// write code here

	// 有序合并
	dummy := &dto2.ListNode{} // 定义一个虚拟节点
	cur := dummy

	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val <= pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else if pHead1.Val > pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}

	// 边界处理
	if pHead1 != nil {
		cur.Next = pHead1
	}
	if pHead2 != nil {
		cur.Next = pHead2
	}
	return dummy.Next
}
