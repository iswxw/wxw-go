/*
@Time: 2022/7/23 15:34
@Author: wxw
@File: BM6_HasCycle
*/
package list

/**
 * 判断链表是否有环
 * @param head ListNode类
 * @return bool布尔型
 */
func hasCycle(head *ListNode) bool {
	// write code here
	// 判断边界
	if head == nil || head.Next == nil {
		return false
	}
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}
