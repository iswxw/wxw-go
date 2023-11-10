/*
@Time: 2023/9/2 22:10
@Author: wxw
@File: BM13_isPail
*/
package _1_list

import (
	"src/com.wxw/project_actual/src/05_algorithm/common/dto"
	"testing"
)

// 题目定义：给定一个链表，请判断该链表是否为回文结构。 回文是指该字符串正序逆序完全一致。
// 方法一：基于数组
// 方法二：双指针
func TestBM13(t *testing.T) {

}

// 基于双指针实现
// 判断链表是否为回文结构
func isPail(head *dto.ListNode) bool {
	// write code here

	return true
}

// 基于数组实现
// 判断链表是否为回文结构 :[1 2 3 4 5 6]
func isPail01(head *dto.ListNode) bool {
	// write code here
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}
