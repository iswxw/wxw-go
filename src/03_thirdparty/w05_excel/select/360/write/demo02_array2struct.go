// @Time : 2022/8/18 11:34
// @Author : xiaoweiwei
// @File : write

package write

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mitchellh/mapstructure"
	"log"
	"os"
	dto2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/dto"
	lib2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/lib"
)

// 相关资料
// https://github.com/liangzibo/go-excel

func Array2Struct() {
	xlsx, err := excelize.OpenFile(GetPath("tmp/test.xlsx"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get all the rows in a sheet.
	rows := xlsx.GetRows("Sheet1")

	//结果在  arr 中
	var arr []dto2.ExcelTest
	err = lib2.NewExcelStructDefault().SetPointerStruct(&dto2.ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
		var ptr dto2.ExcelTest
		// map 转 结构体
		if err2 := mapstructure.Decode(maps, &ptr); err2 != nil {
			return err2
		}
		arr = append(arr, ptr)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("arr1 = ", arr)

	//rows 为[][]string 类型
	//结果在  arr 中
	var arr2 []dto2.ExcelTest
	//StartRow 开始行,索引从 0开始
	//IndexMax  索引最大行,如果 结构体中的 index 大于配置的,那么使用结构体中的
	err = lib2.NewExcelStruct(1, 10).SetPointerStruct(&dto2.ExcelTest{}).RowsAllProcess(rows, func(maps map[string]interface{}) error {
		var ptr dto2.ExcelTest
		// map 转 结构体
		if err2 := mapstructure.Decode(maps, &ptr); err2 != nil {
			return err2
		}
		arr2 = append(arr2, ptr)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("arr2 = ", arr2)
}
