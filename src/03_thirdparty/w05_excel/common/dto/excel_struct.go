// @Time : 2022/8/18 11:17
// @Author : xiaoweiwei
// @File : excel_struct

package dto

import "time"

type ParamWriteExcel struct {
	Headers      []string           `json:"headers"`
	ColumnWidths map[string]float64 `json:"column_widths"`
	FileData     [][]interface{}    `json:"file_data"`
}

type ExcelTest struct {
	Name     string    `json:"name" name:"名称" index:"0"`
	Age      int64     `json:"age" name:"年龄" index:"1"`
	Score    int64     `json:"score" name:"分数" index:"2"`
	Birthday time.Time `json:"birthday" name:"生日" index:"4"`
}
