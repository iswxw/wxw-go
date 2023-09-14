/*
@Time: 2023/9/14 23:13
@Author: wxw
@File: demo03_share_test
*/
package demo

import (
	"fmt"
	"testing"
)

/**
* 切片共享存储空间
 */
func TestSliceShareMemory(t *testing.T) {
	slice1 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	Q2 := slice1[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	// [4 5 6] 3 9

	Q3 := slice1[5:8]
	t.Log(Q3, len(Q3), cap(Q3))
	// [6 7 8] 3 7

	Q3[0] = "Unkown"
	t.Log(Q2, Q3)
	// [4 5 Unkown] [Unkown 7 8]

	a := []int{1, 2, 3, 4, 5}
	shadow := a[1:3]
	t.Log(shadow, a)
	// [2 3] [1 2 3 4 5]

	shadow = append(shadow, 100)
	// 会修改指向数组的所有切片
	t.Log(shadow, a)
	// [2 3 100] [1 2 3 100 5]
}

// 多个切片如果共享同一个底层数组，这种情况下，对其中一个切片或者底层数组的更改，会影响到其他切片。
func TestSliceChangeValue(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */ // 左开右闭
	s1 := slice[2:5]

	// 语法说明：https://blog.csdn.net/qq_45859826/article/details/132587061
	s2 := s1[2:6:7]

	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
