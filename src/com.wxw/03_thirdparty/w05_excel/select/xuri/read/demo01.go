/*
@Time : 2022/1/25 23:15
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"src/com.wxw/project_actual/src/com.wxw/03_thirdparty/w05_excel/common/util"
)

func main() {
	f, err := excelize.OpenFile(util.GetPath("books.xlsx"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(cell)
}
