/*
@Time: 2022/6/3 17:25
@Author: wxw
@File: BM2_reverseBetween
*/
package _1_list

import (
	"fmt"
	"testing"
)

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func TestName2(t *testing.T) {
	head := initList2([]int{1, 2, 3, 4, 5})
	fmt.Println(show2(head))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode2, m int, n int) *ListNode2 {
	// write code here

	return nil
}

// 初始化一个链表
func initList2(values []int) *ListNode2 {
	var head, cur *ListNode2
	for _, v := range values {
		node := &ListNode2{Val: v}
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

// 打印链表
func show2(head *ListNode2) []int {
	if head == nil {
		return []int{}
	}
	results := make([]int, 0)
	cur := head
	for cur != nil {
		results = append(results, cur.Val)
		cur = cur.Next
	}
	return results
}
