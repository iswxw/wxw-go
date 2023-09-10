/*
@Time: 2023/9/10 22:42
@Author: wxw
@File: BM19_findPeakElement_test
*/
package main

import "testing"

// 寻找峰值
// 描述：给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。
//  1. 峰值元素是指其值严格大于左右相邻值的元素。严格大于即不能有等于
func TestBM19(t *testing.T) {

}

// 特征元素取值
func findPeak(nums []int) int {
	n := len(nums)
	up := false

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = true
		} else if up {
			return i - 1
		}
	}

	if !up {
		return 0
	}

	return n - 1
}

// 二分法思想 logN
func findPeakElement(nums []int) int {
	// write code here
	l, r := 0, len(nums)-1

	// 关键思想：下坡的时候可能找到波峰，但是可能找不到，一直向下走的
	// 上坡的时候一定能找到波峰，因为题目给出的是nums[-1] = nums[n] = -∞

	for l < r {
		mid := (r-l)>>1 + l // 防止相加发生溢出
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			//这里是右边的路是上坡路
			l = mid + 1
		}
	}
	return r
}

// 暴力破解法,有些case过不了
func findPeakElement01(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] && nums[i+1] > nums[i+2] {
			return i + 1
		}
	}
	return 0
}
