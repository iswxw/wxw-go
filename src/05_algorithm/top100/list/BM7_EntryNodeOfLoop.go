/*
@Time: 2022/7/23 15:41
@Author: wxw
@File: BM7_EntryNodeOfLoop
*/
package list

// 链表中环的入口节点
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	// 判断边界
	if pHead == nil {
		return pHead
	}

	p1, p2 := pHead, pHead
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			break
		}
	}
	if p2 == nil || p2.Next == nil {
		return nil
	}
	p1 = pHead
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
