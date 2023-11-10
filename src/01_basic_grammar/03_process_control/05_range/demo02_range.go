/*
@Time : 2022/5/1 20:53
@Author : weixiaowei
@File : demo02_range
*/
package main

import "fmt"

func main() {
	funcSlice() // [1 2 3 7 5 6]

	// funcArray() // [7 3 5 7 9 11]

}

func funcArray() {
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
}

func funcSlice() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
}
