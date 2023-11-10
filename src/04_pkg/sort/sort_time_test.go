// @Time : 2023/8/14 17:56
// @Author : xiaoweiwei
// @File : sort_time_test

package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestHello(t *testing.T) {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5 6]
}
