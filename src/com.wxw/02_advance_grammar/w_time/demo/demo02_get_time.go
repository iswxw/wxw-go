/*
@Time : 2022/5/5 11:42
@Author : weixiaowei
@File : demo02_get_time
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// 打印时间戳
	fmt.Println(time.Now().Unix())

	// 格式化字符串时间  // 这是个奇葩,必须是这个时间点, 据说是go诞生之日
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// 时间戳转str格式化时间
	strTime := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
	fmt.Println(strTime)

	// str格式化时间转时间戳
	theTime := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	unixTime := theTime.Unix()
	fmt.Println(unixTime)

	// 还有一种方法,使用time.Parse
	thtime, err := time.Parse("2006-01-02 15:04:05", "2014-01-08 09:04:41")
	if err == nil {
		untime := thtime.Unix()
		fmt.Println(untime)
	}

}
