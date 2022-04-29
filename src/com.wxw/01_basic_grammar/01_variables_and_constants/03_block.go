/*
@Time : 2022/4/29 10:40
@Author : weixiaowei
@File : 03_block
*/
package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
}
