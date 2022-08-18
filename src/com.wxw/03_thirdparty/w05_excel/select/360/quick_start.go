// @Time : 2022/8/18 10:53
// @Author : xiaoweiwei
// @File : quick_start

package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"src/com.wxw/project_actual/src/com.wxw/03_thirdparty/w05_excel/common/util"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs(util.GetPath("src/com.wxw/03_thirdparty/w05_excel/tmp/test.xlsx")); err != nil {
		fmt.Println(err)
	}
}
