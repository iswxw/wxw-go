/*
@Time : 2022/1/24 00:17
@Author : weixiaowei
@File : demo02_batch
*/
package main

import (
	"fmt"
	"framework/w_excel/excelize/common/dto"
	"time"
)

func main() {

	//开始时间
	startTime := time.Now().Unix()
	//创建表格 -创建表格数据 +  创建表格
	createExample()
	//结束时间
	endTime := time.Now().Unix()

	fmt.Println("开始时间： " + fmt.Sprint(startTime))
	fmt.Println("结束时间： " + fmt.Sprint(endTime))
	fmt.Println("消费时间： " + fmt.Sprint(endTime-startTime))

}

/* *
 * 创建表格 示例
 *
 * Desc：
 *    本例中，程序将创建一个10w行、30列的表格用于测试， 本次创建消费时间334s
 **/
func createExample() {
	_sheetData := dto.SheetStruct{}
	_sheetData.SheetName = "Sheet1"

	_sheetList := []dto.SheetStruct{}
	for i := 0; i < 10000; i++ {
		//如果是第一行，则添加行说明
		if i == 0 {
			_lineDesc := dto.LineData{}
			for des := 0; des < 30; des++ {
				_lineDesc.Data = append(_lineDesc.Data, "列."+fmt.Sprint(des+1))
			}
			_sheetData.List = append(_sheetData.List, _lineDesc)
		}
		//拼接line
		_line := dto.LineData{}
		for j := 0; j < 30; j++ {
			//拼接一行中数据
			_line.Data = append(_line.Data, "test-"+fmt.Sprint(i+1)+"-"+fmt.Sprint(j+1))
		}
		_sheetData.List = append(_sheetData.List, _line)
	}
	_sheetList = append(_sheetList, _sheetData)
	dto.CreateXlsxSheet("test-001", _sheetList)

}
