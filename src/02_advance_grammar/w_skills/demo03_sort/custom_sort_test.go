// @Time : 2023/9/6 11:15
// @Author : xiaoweiwei
// @File : custom_sort_test

package demo03_sort

import (
	"fmt"
	"sort"
	"testing"
)

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSortStruct(t *testing.T) {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)
}
