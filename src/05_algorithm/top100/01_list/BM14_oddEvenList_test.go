/*
@Time: 2023/9/2 22:26
@Author: wxw
@File: BM14_oddEvenList_test
*/

// 给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
package _1_list

import (
	"fmt"
	"src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

func TestBM14(t *testing.T) {
	node := init14([]int{1, 2, 3, 4, 5})
	fmt.Println(show14(node))

	evenList := oddEvenList(node)
	fmt.Println(show14(evenList))

}

// 复杂度分析
//  - 时间复杂度：O(n)，其中 n 是链表的节点数。需要遍历链表中的每个节点，并更新指针。
//  - 空间复杂度：O(1)。只需要维护有限的指针。
func oddEvenList(head *dto.ListNode) *dto.ListNode {
	// write code here
	if head == nil {
		return head
	}
	evenHead := head.Next // 偶数头结点
	odd, even := head, evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// 内部方法
func init14(values []int) *dto.ListNode {
	var head, cur *dto.ListNode
	for _, value := range values {
		tmp := &dto.ListNode{
			Val: value,
		}
		if head == nil {
			head = tmp
			cur = head
			continue
		} else if cur != nil {
			cur.Next = tmp
			cur = cur.Next
		}
	}
	return head
}

func show14(head *dto.ListNode) []int {
	cur, results := head, make([]int, 0)
	for cur != nil {
		results = append(results, cur.Val)
		cur = cur.Next
	}
	return results
}
