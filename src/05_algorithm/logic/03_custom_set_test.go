// @Time : 2023/5/31 20:10
// @Author : xiaoweiwei
// @File : 03_custom_set_test

package logic

import (
	"fmt"
	"testing"
)

// 自定义实现set集合,详见：https://www.yuque.com/fcant/go/sakaw6
func TestCustomSet(t *testing.T) {
	set := NewSet()
	set.Add("hello")
	set.Add("world")
	fmt.Println(set.Contains("hello"))

	set.Remove("hello")
	fmt.Println(set.Contains("hello"))
}

// 定义了一个保存 string 类型的 Set集合
type Set map[string]struct{}

// 添加一个元素
func (s Set) Add(key string) {
	s[key] = struct{}{}
}

// 移除一个元素
func (s Set) Remove(key string) {
	delete(s, key)
}

// 是否包含一个元素
func (s Set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}

// 初始化
func NewSet() Set {
	s := make(Set)
	return s
}
