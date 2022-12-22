/*
@Time: 2022/12/22 10:27
@Author: wxw
@File: 27_pprof_visual
*/
package main

import "testing"

const url = "https://github.com/iswxw"

func TestAdd(t *testing.T) {
	s := AddList(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(url)
	}
}

var dataList []string

func AddList(str string) string {
	data := []byte(str)
	sData := string(data)
	dataList = append(dataList, sData)
	return sData
}
