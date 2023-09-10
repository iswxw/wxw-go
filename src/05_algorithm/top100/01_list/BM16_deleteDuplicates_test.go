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
// 复杂度分析：时间复杂度：O(n) 、空间复杂度O(1)
func deleteDuplicate01(head *dto.ListNode) *dto.ListNode {
	if head == nil {
		return head
	}

	// 开一个虚拟头结点
	dummy := &dto.ListNode{}
	dummy.Next = head
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {

		// 相邻两个节点值相同
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

// 方法三：哈希法，对于无序的也可以使用
func deleteDuplicate02(head *dto.ListNode) *dto.ListNode {

	// 排除空链表
	if head == nil {
		return head
	}

	// 定义一个map记录值出现的次数
	hashMap := make(map[int]int, 0)
	cur := head
	// 遍历统计每个节点值出现的次数
	for cur != nil {
		if _, ok := hashMap[cur.Val]; ok {
			hashMap[cur.Val] = hashMap[cur.Val] + 1
		} else {
			hashMap[cur.Val] = 1
		}
		cur = cur.Next
	}

	// 开一个虚拟头结点
	dummy := &dto.ListNode{}
	dummy.Next = head
	temp := dummy
	for temp.Next != nil {
		// 如果节点值计数不为1
		if hashMap[temp.Next.Val] != 1 {
			// 删去该节点
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}
	//去掉表头
	return dummy.Next
}
