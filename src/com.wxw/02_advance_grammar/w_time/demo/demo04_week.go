// @Time : 2022/8/25 14:29
// @Author : xiaoweiwei
// @File : demo04_week

package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	week := now.Weekday()
	offset := int(time.Monday - week)
	if offset > 0 {
		offset = -6
	}
	fmt.Println(offset)

	//周
	StartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	EndDates := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, time.Local)
	fmt.Println("周一", StartDate.AddDate(0, 0, offset))            //周一
	fmt.Println("周末", EndDates.AddDate(0, 0, int((week+1)-week))) //周末
}
