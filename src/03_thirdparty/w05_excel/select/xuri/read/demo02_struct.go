/*
@Time : 2022/1/27 19:54
@Author : weixiaowei
@File : demo02_struct
*/
package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	dto2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/dto"
	util2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/util"
)

// 来源 https://github.com/liangzibo/go-excel

func main() {
	xlsx, err := excelize.OpenFile(util2.GetPath("books.xlsx"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get all the rows in a sheet.
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var demoStruct dto2.ExcelTest
	// map 转 结构体
	if err2 := mapstructure.Decode(rows, &demoStruct); err2 != nil {
		log.Println(err2)
	}
	fmt.Println(demoStruct)
}
