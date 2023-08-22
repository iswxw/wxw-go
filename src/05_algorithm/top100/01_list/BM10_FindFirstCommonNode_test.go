/*
@Time: 2023/8/22 23:05
@Author: wxw
@File: BM10_FindFirstCommonNode_test
*/
package _1_list

import "src/com.wxw/project_actual/src/05_algorithm/common/dto"

// 说明：
// 因为两个指针，同样的速度，走完同样长度（链表1+链表2），不管两条链表有无相同节点，都能够到达同时到达终点。
//（N1最后肯定能到达链表2的终点，N2肯定能到达链表1的终点）。

// FindFirstCommonNode 双指针
func FindFirstCommonNode(pHead1 *dto.ListNode, pHead2 *dto.ListNode) *dto.ListNode {
	// write code here

	p1, p2 := pHead1, pHead2

	for p1 != p2 {
		if p1 == nil {
			p1 = pHead2
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = pHead1
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
