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
func Find( target int ,  array [][]int ) bool {
	// write code here
	row, col := 0, len(array[0])-1
	for row <= len(array)-1 && col >= 0 {
		if target == array[row][col] {
			return true
		} else if target > array[row][col] {
			row++
		} else {
			col--
		}
	}
	return false
}
