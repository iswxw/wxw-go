// @Time : 2023/7/25 14:26
// @Author : xiaoweiwei
// @File : BM1_reverse_list_test

package _1_list

import (
	"fmt"
	"testing"
)

// ListNode1 定义节点
type ListNode1 struct {
	Val  int        `json:"val"`
	Next *ListNode1 `json:"next"`
}

func TestHello(t *testing.T) {
	head := initList([]int{1, 2, 3, 4, 5})
	fmt.Println("1.查看原始链表", show1(head))

	head = reverse1(head)
	fmt.Println("2.指针法：翻转链表", show1(head))

	fmt.Println("3.递归法：翻转链表", show1(reverse2(head)))
}

// 工具类
// newListNode1 新建节点
func newListNode1(value int) *ListNode1 {
	return &ListNode1{Val: value}
}
func initList(values []int) *ListNode1 {
	var head, cur *ListNode1
	for _, v := range values {
		node := newListNode1(v)
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
func show1(head *ListNode1) []int {
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

// 算法实现
// 方式一：迭代实现
func reverse1(pHead *ListNode1) *ListNode1 {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	// 1,2,3
	var newHead *ListNode1
	for pHead != nil {
		pNext := pHead.Next  // 保留未反转的链表
		pHead.Next = newHead // 节点反转
		newHead = pHead      // 更新已反转的节点
		pHead = pNext        // 更新当前节点
	}
	return newHead
}

// 方式二：递归实现
func reverse2(pHead *ListNode1) *ListNode1 {
	// 递归结束条件：只有一个节点（或者头结点为空，直接不进行递归直接返回）
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	newHead := reverse2(pHead.Next)

	// 让head的下一个节点反过来指向head
	pHead.Next.Next = pHead

	// 将head的next置空
	pHead.Next = nil

	return newHead
}
