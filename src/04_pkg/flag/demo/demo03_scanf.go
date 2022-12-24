/*
@Time : 2022/4/28 16:54
@Author : weixiaowei
@File : demo03_scanf
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var age, name string
	// Scanf要求 严格按照输出格式进行输入
	fmt.Scanf("%s,%s", &name, &age)
	fmt.Printf("my name is %s,i am %s years old .\n", name, age)
	fmt.Printf("%s\n", os.Args)
	//这种方式需要 单独输入参数.
}

// 输出提示
// $ go run demo03_scanf.go
//张三 18
//my name is 张三,i am  years old .
//[/var/folders/vd/ms9cv4zj187g_wd6bzt4jq2r0000ks/T/go-build1002993775/b001/exe/demo03_scanf]
