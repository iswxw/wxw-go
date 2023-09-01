// @Time : 2023/8/29 16:46
// @Author : xiaoweiwei
// @File : BM12_sortInList 单链表排序

package _1_list

import (
	"fmt"
	"sort"
	"src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

// 题目：单链表排序
// 方法：1.辅助数组、2.归并排序
// 复杂度分析：时间复杂度O(NlogN)，空间复杂度：方法一O(N)、方法二O(1)
// 材料：https://zhuanlan.zhihu.com/p/568346489
func TestBM12(t *testing.T) {
	bm12 := bm12Init([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(bm12Show(bm12))

	// 方法一：基于数据给链表排序
	bm121 := bm12Init([]int{7, 3, 1, 5, 8, 6})
	sort12 := bm12Sort01(bm121)
	fmt.Println(bm12Show(sort12))

	// 方法二：基于归并算法排序

}

func TestBM121(t *testing.T) {
	node := &dto.ListNode{}
	fmt.Println(node)
	fmt.Println(node == nil)
	fmt.Printf("%#v \n", node)
	fmt.Printf("%v \n", node)
}

// 方法二：通过递归实现链表的归并排序
// 时间复杂度O(NlogN)：N表示链表结点数量，二分归并算法O(NlogN)
// 空间复杂度O(1)：仅使用常数级变量空间
func bm12Sort(head *dto.ListNode) *dto.ListNode {

	// 判断边界
	if head == nil {
		return head
	}

	// 中间节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	temp := slow.Next
	slow.Next = nil

	// 递归左右两边进行排序
	l := bm12Sort(head)
	r := bm12Sort(temp)

	var h, cur *dto.ListNode
	for l != nil && r != nil {
		// left  right链表循环对比
		if l.Val < r.Val {
			h.Next = l
			l = l.Next
		} else {
			h.Next = r
			r = r.Next
		}
		h = h.Next
	}
	return cur
}

// 方法一：数组辅助排序
// 时间复杂度O(NlogN)：N表示链表结点数量，遍历链表O(N)，数组排序(NlogN)，遍历数组O(N)
// 空间复杂度O(N)：使用额外数组占用空间O(N)
func bm12Sort01(head *dto.ListNode) *dto.ListNode {
	//head = bm12Init([]int{1, 2, 3, 4, 5, 6})

	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	sort.Ints(values)

	var cur, newHead *dto.ListNode
	for _, value := range values {
		temp := &dto.ListNode{
			Val: value,
		}
		if newHead == nil {
			newHead = temp
			cur = newHead
			continue
		}
		cur.Next = temp
		cur = cur.Next
	}
	return newHead
}

// =================内部方法=======================
func bm12Init(values []int) *dto.ListNode {
	var cur, head *dto.ListNode
	for _, elem := range values {
		node := &dto.ListNode{Val: elem}
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

func bm12Show(head *dto.ListNode) []int {
	cur, results := head, make([]int, 0)
	for cur != nil {
		results = append(results, cur.Val)
		cur = cur.Next
	}
	return results
}
