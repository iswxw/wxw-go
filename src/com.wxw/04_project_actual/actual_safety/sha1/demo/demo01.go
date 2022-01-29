/*
@Time : 2022/1/29 23:37
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	sum := sha1.Sum([]byte("maple"))
	fmt.Println(fmt.Sprintf("%x", sum))
}
