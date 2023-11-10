// @Time : 2023/6/1 15:26
// @Author : xiaoweiwei
// @File : demo01_generic_test

package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenericLearning(t *testing.T) {
	var p1 p[string, struct {
		Name string
		Age  int
	}]
	var p2 p[string, int64]

	fmt.Println(reflect.TypeOf(p1))
	fmt.Println(reflect.TypeOf(p2))
}

type p[T, T1 any] struct {
	A string
	B T
	C T1
}
