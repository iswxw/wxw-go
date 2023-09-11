/*
@Time: 2022/7/24 17:46
@Author: wxw
@File: BM18_two_array_search
*/
package main

/**
 * 二维数组中的查找
[
 [1,2,8,9],
 [2,4,9,12],
 [4,7,10,13],
 [6,8,11,15]
]
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
*/

// 方法一：双指针移动法（右上角）
func Find(target int, array [][]int) bool {
	// write code here
	row, col := 0, len(array[0])-1
	for row <= len(array)-1 && col >= 0 {

		if target == array[row][col] {
			return true
		} else if target > array[row][col] {
			//如果目标值大于当前位置的值，当前位置右移
			row++
		} else {
			//如果目标值小于当前位置的值，当前位置上移
			col--
		}
	}
	return false
}

// 方法二：暴力破解法
func Find01(target int, array [][]int) bool {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if target == array[i][j] {
				return true
			}
		}
	}
	return false
}
