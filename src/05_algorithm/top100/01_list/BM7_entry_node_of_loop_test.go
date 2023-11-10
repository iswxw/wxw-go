/*
@Time: 2022/7/23 15:41
@Author: wxw
@File: BM7_EntryNodeOfLoop BM7 链表中环的入口结点
*/
package _1_list

import (
	dto2 "src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func TestBM7(t *testing.T) {

}

// EntryNodeOfLoop  判断环的入口节点： 快慢指针法
func EntryNodeOfLoop(pHead *dto2.ListNode) *dto2.ListNode {
	// 边界判断
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	p1, p2 := pHead, pHead
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next

		// 相交说明有环
		if p1 == p2 {
			// p1 相交点 p2-头结点
			p2 = pHead
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			return p1
		}
	}
	return nil
}
