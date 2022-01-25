/*
@Time : 2022/1/26 01:03
@Author : weixiaowei
@File : demo02
*/
package main

import (
	"fmt"
	"framework/w_excel/excelize/common/util"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"time"
)

func main() {

	header := []interface{}{"编号", "姓名", "年龄"}
	data := [][]interface{}{
		{"1", "北京", 13},
		{"2", "天津", 18},
	}
	ExcelExportData360(header, data, "Sheet1")
}

func ExcelExportData360(header []interface{}, data [][]interface{}, sheetName string) {

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet(sheetName)

	if err := f.SetSheetRow(sheetName, "A1", &header); err != nil {
		log.Println("err1= ", err)
		return
	}
	axi := 1
	for _, obj := range data {
		axi = axi + 1
		if err := f.SetSheetRow(sheetName, "A"+strconv.Itoa(axi), &obj); err != nil {
			log.Println("err2= ", err)
			return
		}
	}
	f.SetActiveSheet(index)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	// Save xlsx file by the given path.
	if err := f.SaveAs(util.GetPath(timeStr + ".xlsx")); err != nil {
		fmt.Println(err)
	}
}
