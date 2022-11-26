/*
@Time : 2022/1/26 01:21
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	util2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/util"
)

func main() {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet2")
	// 设置单元格的值
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs(util2.GetPath("Book1.xlsx")); err != nil {
		fmt.Println(err)
	}
}
