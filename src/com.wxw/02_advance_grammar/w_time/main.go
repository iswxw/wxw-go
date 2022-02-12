/*
@Time : 2022/1/25 15:52
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01"))
	fmt.Println(time.Now().Format("20060102150405"))
}
