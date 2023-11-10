/*
@Time: 2021/12/5 0:41
@Author: wxw
@File: unit_test
*/
package main

import (
	"reflect"
	"testing"
)

//测试用例1：以字符分割
func TestSplit(t *testing.T) {
	got := NewSplit("123N456", "N")
	want := []string{"123", "456"}
	//DeepEqual比较底层数组
	if !reflect.DeepEqual(got, want) {
		//如果got和want不一致说明你写得代码有问题
		t.Errorf("The values of %v is not %v\n", got, want)
	}
}

//测试用例2：以标点符号分割
func TestPunctuationSplit(t *testing.T) {
	got := NewSplit("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.FailNow() //出错就stop别往下测了！
	}

}
