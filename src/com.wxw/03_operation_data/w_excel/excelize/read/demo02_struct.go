/*
@Time : 2022/1/27 19:54
@Author : weixiaowei
@File : demo02_struct
*/
package main

import (
	"fmt"
	"framework/w_excel/excelize/common/dto"
	"framework/w_excel/excelize/common/util"
	"github.com/mitchellh/mapstructure"
	"github.com/xuri/excelize/v2"
	"os"
)

// 来源 https://github.com/liangzibo/go-excel

func main() {
	xlsx, err := excelize.OpenFile(util.GetPath("books.xlsx"))
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

	var demoStruct dto.ExcelTest
	// map 转 结构体
	if err2 := mapstructure.Decode(rows, &demoStruct); err2 != nil {

	}
	fmt.Println(demoStruct)
}
