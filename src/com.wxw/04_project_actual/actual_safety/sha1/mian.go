/*
@Time : 2022/1/24 15:43
@Author : weixiaowei
@File : mian
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
