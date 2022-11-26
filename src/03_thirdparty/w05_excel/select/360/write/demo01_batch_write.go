// @Time : 2022/8/18 11:34
// @Author : xiaoweiwei
// @File : write

package write

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	util2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/util"
	"strconv"
)

func batchWrite() {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")

	axi := 1
	for _, obj := range demoAll {
		axi = axi + 1
		f.SetSheetRow("Sheet1", "A"+strconv.Itoa(axi), &obj)
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 根据指定路径保存文件
	if err := f.SaveAs(util2.GetPath("src/com.wxw/03_thirdparty/w05_excel/tmp/test.xlsx")); err != nil {
		fmt.Println(err)
	}
}

var demoAll = [][]string{
	{"赵三",
		"30",
		"100",
		"1970-01-01 12:50:01",
		"1980-11-21 15:20:01"},
	{"赵六",
		"40",
		"30",
		"1970-01-01 12:50:01",
		"1980-11-21 15:20:01"},
}
