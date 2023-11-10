package main

import (
	"container/ring"
	"fmt"
)

// ring 环形容器集合： 环可以理解为前后节点相连的链表
func main() {
	r := ring.New(3)
	fmt.Println("最开始的环属性信息：", r)
}
