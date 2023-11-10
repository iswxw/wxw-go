/*
@Time : 2022/4/29 14:53
@Author : weixiaowei
@File : demo02_map
*/
package main

import "fmt"

func main() {

	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}
	fmt.Println("v=", v, "ok=", ok)
}
