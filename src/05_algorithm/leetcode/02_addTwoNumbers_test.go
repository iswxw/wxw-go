/*
@Time: 2023/6/11 10:07
@Author: wxw
@File: 02_addTwoNumbers_test https://leetcode.cn/problems/add-two-numbers/
*/
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：基础数学
// 方法二：递归

// 基础实现
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
func addTwoNumbers01(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur *ListNode
	var bit int

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + bit

		// 分别取个位sum，和进位bit
		sum, bit = sum%10, sum/10

		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			cur = head
		} else {
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
		}
	}
	if bit > 0 {
		cur.Next = &ListNode{Val: bit}
	}

	return head
}

// 递归
func addTwoNumbers02(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

// 返回两个相加链表的头部
func addTwo(l1 *ListNode, l2 *ListNode, bit int) *ListNode {
	if l1 == nil && l2 == nil && bit == 0 {
		return nil
	}
	sum := bit

	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	cur := &ListNode{Val: sum % 10}
	cur.Next = addTwo(l1, l2, sum/10)
	return cur
}
