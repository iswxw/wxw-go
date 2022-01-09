/*
@Time : 2022/1/5 14:06
@Author : wxw
@File : demo_rand
*/
package main

import (
	"fmt"
	"math/bits"
	"math/rand"
	"reflect"
)

func main() {
	val := rand.Int()
	fmt.Println(reflect.TypeOf(val))
	fmt.Println(bits.Len(uint(val)))
	fmt.Println(rand.Uint32())
	fmt.Println(val)
}
