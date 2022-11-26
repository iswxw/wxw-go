/*
@Time : 2022/4/29 14:49
@Author : weixiaowei
@File : demo01_sort
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// To create a map as input
	m := make(map[string]int)
	m["tom"] = 2
	m["jame"] = 4
	m["amy"] = 5

	// To store the keys in slice in sorted order
	var strs []string
	for k := range m {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	// To perform the opertion you want
	for _, k := range strs {
		fmt.Printf("%s\t%d\n", k, m[k])
	}
}
