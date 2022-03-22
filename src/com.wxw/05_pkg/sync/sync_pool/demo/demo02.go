/*
@Time : 2022/3/22 17:02
@Author : weixiaowei
@File : demo02
*/
package main

import (
	"fmt"
	"sync"
)

// 定义一个池
var pool *sync.Pool

// 声明一个对象
type Person struct {
	Name string
}

// 初始化 对象池
func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return &Person{}
		},
	}
}

func main() {
	initPool()

	p := pool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	pool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person)) // 断言：强制转换类型
}
