/*
@Time : 2022/4/28 15:22
@Author : weixiaowei
@File : demo02_flag_type
*/
package main

import "fmt"

// https://www.cnblogs.com/aaronthon/p/10883675.html
func main() {
	var (
		name    string
		age     int
		married bool
	)
	// fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。
	fmt.Scan(&name, &age, &married)

	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// 输出结果
// case didi$ go run demo02_scan.go
//aq
//12
//12
//扫描结果 name:aq age:12 married:true
