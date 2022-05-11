/*
@Time : 2022/5/5 11:34
@Author : weixiaowei
@File : test
*/
package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	m := map[string]interface{}{
		"2": 2,
	}

	valueStr := m["2"]
	fmt.Println(valueStr)

	value := valueStr.(int)

	fmt.Println(cast.ToInt(valueStr))
	fmt.Println(value)
}
