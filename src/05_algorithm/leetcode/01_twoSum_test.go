/*
@Time : 2022/1/11 12:52
@Author : wxw
@File : 01_twoSum
@Link : https://leetcode.cn/problems/two-sum/
*/
package main

import (
	"fmt"
	"testing"
)

// 题目描述： 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

func TestTwoSum(t *testing.T) {
	// 定义一个数组
	nums := []int{1, 2, 3, 4}
	arrayIndex := twoSum(nums, 4)
	fmt.Println("arrayIndex = ", arrayIndex)
}

// 方法一：暴力遍历  O(n2)
// 方法二：hash分析 O（n）

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)
	for i, num := range nums {
		if v, ok := hashMap[target-num]; ok {
			return []int{v, i}
		}
		hashMap[num] = i
	}
	return nil
}
