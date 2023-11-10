/*
@Time: 2023/8/22 23:36
@Author: wxw
@File: BM11_addInList_test
*/
package _1_list

import (
	"fmt"
	"src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

// 方法：基础数据<递归
// 两个逆序链表相加，返回新的链表表示它们的和。
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

// 题目描述：给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
func TestBM11(t *testing.T) {
	l1 := initB11([]int{2, 4, 3})
	l2 := initB11([]int{5, 6, 4})
	addList := addInList(l1, l2)
	//addList := addInList01(l1, l2)
	fmt.Println(showB11(addList))
}

func TestBM111(t *testing.T) {
	fmt.Printf("取十位值：%d,取余数：%d \n", 16/10, 16%10)
}

// addInList 两个链表求和：基础数学
func addInList(head1 *dto.ListNode, head2 *dto.ListNode) *dto.ListNode {
	// write code here

	// 定义新链表的信息
	var head, cur *dto.ListNode
	var bit int // 进位值

	// 计算加法
	for head1 != nil || head2 != nil {
		v1, v2 := 0, 0 // 暂存两个节点的数值

		if head1 != nil {
			v1 = head1.Val
			head1 = head1.Next
		}

		if head2 != nil {
			v2 = head2.Val
			head2 = head2.Next
		}

		// 求和取值
		sum := v1 + v2 + bit
		sum, bit = sum%10, sum/10 // 分别取相加后的 个位和进位（十位）

		if head == nil {
			head = &dto.ListNode{Val: sum}
			cur = head
		} else {
			cur.Next = &dto.ListNode{Val: sum}
			cur = cur.Next
		}
	}

	if bit > 0 {
		cur.Next = &dto.ListNode{Val: bit}
	}
	return head
}

// addInList 两个链表求和:递归
func addInList01(head1 *dto.ListNode, head2 *dto.ListNode) *dto.ListNode {
	// write code here
	return addList(head1, head2, 0)
}

// addList
func addList(head1 *dto.ListNode, head2 *dto.ListNode, bit int) *dto.ListNode {

	// 边界判断
	if head1 == nil && head2 == nil && bit == 0 {
		return nil
	}

	sum := bit

	// 处理分析
	if head1 != nil {
		sum += head1.Val
		head1 = head1.Next
	}

	if head2 != nil {
		sum += head2.Val
		head2 = head2.Next
	}

	cur := &dto.ListNode{Val: sum % 10}
	cur.Next = addList(head1, head2, sum/10)
	// 结果记录
	return cur
}

// 初始化链表
func initB11(values []int) *dto.ListNode {
	var cur, head *dto.ListNode
	for _, value := range values {
		node := &dto.ListNode{Val: value}
		if head == nil {
			head = node
			cur = head
			continue
		}
		cur.Next = node
		cur = cur.Next
	}
	return head
}

func showB11(head *dto.ListNode) []int {
	cur := head
	results := make([]int, 0)
	for cur != nil {
		results = append(results, cur.Val)
		cur = cur.Next
	}
	return results
}
