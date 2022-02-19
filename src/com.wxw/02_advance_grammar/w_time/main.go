/*
@Time : 2022/1/25 15:52
@Author : weixiaowei
@File : main
*/
package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01"))
	fmt.Println(time.Now().Format("20060102150405"))
	//date, _ := cast.StringToDate("2021-04-16 00:00:00")
	location, _ := cast.StringToDateInDefaultLocation("2021-04-16 00:00:00", time.Local)
	fmt.Println(location)

	fmt.Println(cast.ToInt(""))
}
