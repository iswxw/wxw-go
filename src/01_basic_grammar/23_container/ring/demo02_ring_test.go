// @Time : 2023/5/30 17:34
// @Author : xiaoweiwei
// @File : demo02_ring_test

package ring

import (
	"container/ring"
	"fmt"
	"testing"
)

// TestRingHelloWorld ring 是一种环形链表结构
func TestRingHelloWorld(t *testing.T) {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})
}
