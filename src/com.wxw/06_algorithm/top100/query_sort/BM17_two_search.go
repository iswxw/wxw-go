/*
@Time: 2022/7/24 16:53
@Author: wxw
@File: BM17_two_search
*/
package main


/**
 * 二分法查找
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func search( nums []int ,  target int ) int {
	// write code here
	start,end := 0,len(nums)-1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if target > nums[mid] {
			start = mid +1
		}else if target < nums[mid] {
			end = mid-1
		} else {
			return mid
		}
	}
	return -1
}



