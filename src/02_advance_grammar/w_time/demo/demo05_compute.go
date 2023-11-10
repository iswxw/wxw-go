// @Time : 2022/11/15 16:39
// @Author : xiaoweiwei
// @File : demo05_duration

package main

import (
	"fmt"
	"time"
)

// main 参考：https://www.zhangbj.com/p/652.html
func main() {
	t := time.Now()
	addOneHour := t.Add(time.Hour)
	addTwoHour := t.Add(2 * time.Hour)
	fmt.Println("增加1小时：", addOneHour)
	fmt.Println("增加2小时：", addTwoHour)

	subTwoHour := t.Add(-2 * time.Hour)
	fmt.Println("减去2小时：", subTwoHour)

	addDate := t.AddDate(1, 0, 0)
	fmt.Println("增加1年：", addDate) // 2021-10-24 22:10:53.328973 +0800 CST

	subDate := t.AddDate(-1, 0, 0)
	fmt.Println("减去1年：", subDate) // 2019-10-24 22:10:53.328973 +0800 CST

	before := t.Before(t.Add(time.Hour))
	fmt.Println("before：", before)

	after := t.After(t.Add(time.Hour))
	fmt.Println("after：", after)
}
