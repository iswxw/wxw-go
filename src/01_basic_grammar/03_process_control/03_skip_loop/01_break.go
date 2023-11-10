/*
@Time : 2020/11/15 10:17
@Author : wxw
@File : 01_goto
@Software: GoLand
*/
package main

import "fmt"

// 流程控制之跳出for循环，break、continue、goto , 当然这些关键词都可以用 跳转到指定标签的
func main() {
	// 当 i= 5时，跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出for循环
		}
		fmt.Printf("%d 、", i)
	}
	fmt.Println("\n game over")

	// ======================================================

	// 当i=5时，跳过此次for循环（不执行for循环内部的打印语句），继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 结束本次循环跳出下一次循环
		}
		fmt.Printf("%d.", i)
	}
	fmt.Println("\n game over")

	// ===============goto ================
	// goto + lebel 实现跳出多层循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 跳转到指定的标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
breakHere: // 标签
	fmt.Println("跳转到我这个标签啦，done")

}
