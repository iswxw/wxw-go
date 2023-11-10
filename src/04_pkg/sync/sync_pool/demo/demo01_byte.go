/*
@Time: 2021/12/5 22:36
@Author: wxw
@File: demo_loadObject
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 构建一个pool
var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

// https://www.cnblogs.com/sunsky303/p/9706210.html
func main() {
	//defer
	//test.SetGCPercent(test.SetGCPercent(-1))

	// 记录开始构建对象时间
	a := time.Now().Unix() // 返回单位是 秒（s）

	// 构建对象到内存中
	for i := 0; i < 1000000000; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}

	// 对象构建完成时间
	b := time.Now().Unix()

	for j := 0; j < 1000000000; j++ {

		// 获取对象
		obj := bytePool.Get().(*[]byte)
		_ = obj

		// 放入对象
		bytePool.Put(obj)
	}

	// 存取对象
	c := time.Now().Unix()

	fmt.Println("without pool ", b-a, "s")
	fmt.Println("with    pool ", c-b, "s")
}
