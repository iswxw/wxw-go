/*
@Time: 2021/10/12 0:23
@Author: wxw
@File: demo_panic_recover
*/
package main

import "fmt"

func doRecover() {
	fmt.Println("捕获到异常 =>", recover()) //输出: 捕获到异常 => <nil>
}
func main() {
	defer func() {
		doRecover() //注意：这里间接使用函数，在函数中调用了recover()函数，
		// panic 没有恢复,没有捕获到错误信息
	}()
	panic(any("手动抛出异常"))
}
