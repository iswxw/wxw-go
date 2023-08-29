// @Time : 2023/8/29 16:46
// @Author : xiaoweiwei
// @File : BM12_sortInList 单链表排序

package _1_list

import (
	"fmt"
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
}

// 方法二：归并排序
// 时间复杂度O(NlogN)：N表示链表结点数量，二分归并算法O(NlogN)
// 空间复杂度O(1)：仅使用常数级变量空间
func bm12Sort(head *dto.ListNode) *dto.ListNode {

	return nil
}

// 方法一：数组辅助排序
// 时间复杂度O(NlogN)：N表示链表结点数量，遍历链表O(N)，数组排序(NlogN)，遍历数组O(N)
// 空间复杂度O(N)：使用额外数组占用空间O(N)
func bm12Sort01(head *dto.ListNode) *dto.ListNode {

	return nil
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
		cur.Next = &dto.ListNode{Val: elem}
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
