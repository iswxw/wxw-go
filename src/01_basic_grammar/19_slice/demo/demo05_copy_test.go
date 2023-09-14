/*
@Time: 2023/9/14 23:27
@Author: wxw
@File: demo05_copy_test
*/
package demo

import "testing"

/**
 * 深拷贝
 */
func TestSliceDeepCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5, 5)
	// 深拷贝
	copy(slice2, slice1)
	t.Log(slice1, len(slice1), cap(slice1))
	// [1 2 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2))
	// [1 2 3 4 5] 5 5

	slice1[1] = 100
	t.Log(slice1, len(slice1), cap(slice1))
	// [1 100 3 4 5] 5 5

	t.Log(slice2, len(slice2), cap(slice2))
	// [1 2 3 4 5] 5 5
}

/**
 * 浅拷贝
 */
func TestSliceShadowCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	// 浅拷贝（注意 := 对于引用类型是浅拷贝，对于值类型是深拷贝）
	slice2 := slice1
	t.Logf("%p", slice1) // 0xc00001c120
	t.Logf("%p", slice2) // 0xc00001c120

	// 同时改变两个数组，这时就是浅拷贝，未扩容时，修改 slice1 的元素之后，slice2 的元素也会跟着修改
	slice1[0] = 10
	t.Log(slice1, len(slice1), cap(slice1))
	// [10 2 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2))
	// [10 2 3 4 5] 5 5
	// 注意下：扩容后，slice1和slice2不再指向同一个数组，修改 slice1 的元素之后，slice2 的元素不会被修改了
	slice1 = append(slice1, 5, 6, 7, 8)
	slice1[0] = 11
	// 这里可以发现，slice1[0] 被修改为了 11, slice1[0] 还是10
	t.Log(slice1, len(slice1), cap(slice1))
	// [11 2 3 4 5 5 6 7 8] 9 10
	t.Log(slice2, len(slice2), cap(slice2))
	// [10 2 3 4 5] 5 5
}
