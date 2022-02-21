/*
@Time : 2022/2/21 22:10
@Author : weixiaowei
@File : format_time
*/
package util

import (
	"encoding/json"
	"time"
)

type FormatTime time.Time

//TimeFormat 格式化时间
var TimeFormat = "2006-01-02 15:04:05"

func (s *FormatTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	// Ignore null, like in the main JSON package.
	if str == "" || str == `""` {
		return nil
	}
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation(TimeFormat, str[1:len(str)-1], loc)
	if err != nil {
		return err
	}
	*s = FormatTime(t)
	return nil
}

func (s FormatTime) MarshalJSON() ([]byte, error) {
	t := time.Time(s)
	if y := t.Year(); y < 2016 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return json.Marshal("")
	} else {
		str := t.Format(TimeFormat)
		return json.Marshal(str)
	}
}

func (s *FormatTime) Value() time.Time {
	return time.Time(*s)
}
