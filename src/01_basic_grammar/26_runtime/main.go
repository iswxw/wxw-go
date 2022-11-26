/*
@Time: 2022/10/15 18:51
@Author: wxw
@File: main
*/
package main

import "runtime"

func main() {
	println(runtime.GOMAXPROCS(1))
}
