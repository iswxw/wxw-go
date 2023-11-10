// @Time : 2023/6/1 11:16
// @Author : xiaoweiwei
// @File : demo_generic_test

package _0_generic

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index[string](si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index[string](ss, "hello"))
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
