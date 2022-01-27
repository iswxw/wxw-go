/*
@Time : 2022/1/25 23:15
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"fmt"
	"framework/w_excel/excelize/common/util"
	"github.com/xuri/excelize/v2"
	"log"
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
