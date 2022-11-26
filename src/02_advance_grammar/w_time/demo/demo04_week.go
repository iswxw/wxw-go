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
	currentMonday := GetMondayOfWeek(now, "2006-01-02")
	lastMonday, err := GetLastWeekMonday(now, "2006-01-02")
	lastSunday, err := GetLastWeekSunday(now, "2006-01-02")
	fmt.Println(err)
	fmt.Println(currentMonday)
	fmt.Println(lastMonday)
	fmt.Println(lastSunday)
}

//获取本周周一的日期
func GetMondayOfWeek(t time.Time, fmtStr string) (dayStr string) {
	dayObj := GetZeroTime(t)
	if t.Weekday() == time.Monday {
		//修改hour、min、sec = 0后格式化
		dayStr = dayObj.Format(fmtStr)
	} else {
		offset := int(time.Monday - t.Weekday())
		if offset > 0 {
			offset = -6
		}
		dayStr = dayObj.AddDate(0, 0, offset).Format(fmtStr)
	}
	return
}

//获取上周周一日期
func GetLastWeekMonday(t time.Time, fmtStr string) (day string, err error) {
	monday := GetMondayOfWeek(t, fmtStr)
	dayObj, err := time.Parse(fmtStr, monday)
	if err != nil {
		return
	}
	day = dayObj.AddDate(0, 0, -7).Format(fmtStr)
	return
}

//获取上周周日日期
func GetLastWeekSunday(t time.Time, fmtStr string) (day string, err error) {
	monday := GetMondayOfWeek(t, fmtStr)
	dayObj, err := time.Parse(fmtStr, monday)
	if err != nil {
		return
	}
	day = dayObj.AddDate(0, 0, -1).Format(fmtStr)
	return
}

//获取某一天的0点时间
func GetZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
