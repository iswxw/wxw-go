/*
@Time: 2023/6/11 14:50
@Author: wxw
@File: 03_lengthOfLongestSubstring_test
*/
package leetcode

import (
	"testing"
)

// 题目：3. 无重复字符的最长子串
// 思路：滑动窗口或双指针，复杂度O(n)
func TestLengthOfLongestSubstring(t *testing.T) {

}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	r, res := -1, 0

	// i 为左指针，r为右指针
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}

		for r+1 < n && m[s[r+1]] == 0 {
			// 不断地移动右指针
			m[s[r+1]]++
			r++
		}

		// 第 i 到 r 个字符是一个极长的无重复字符子串
		res = max(res, r-i+1)
	}

	return res
}

//结果验证
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
