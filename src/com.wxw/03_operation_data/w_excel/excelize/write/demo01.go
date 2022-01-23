/*
@Time : 2022/1/23 23:51
@Author : weixiaowei
@File : demo01
*/
package main

import (
	"framework/w_excel/excelize/common/util"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	testWriteData01()
}

func testWriteData01() {
	f := excelize.NewFile()
	// 设置单元格的值
	f.SetCellValue("Sheet1", "B1", 100)

	// 按行赋值
	if err := f.SetSheetRow("Sheet1", "A2", &[]interface{}{"1", "Java半颗糖", 2}); err != nil {
		log.Println("err:", err)
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(util.GetPath("Book1.xlsx")); err != nil {
		log.Println("err:", err)
	}
}
