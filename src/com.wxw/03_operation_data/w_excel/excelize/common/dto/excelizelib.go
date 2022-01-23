/*
@Time : 2022/1/24 00:20
@Author : weixiaowei
@File : excelizelib
*/
package dto

import (
	"fmt"
	"framework/w_excel/excelize/common/util"
	"github.com/xuri/excelize/v2"
)

type SheetStruct struct {
	SheetName string
	List      []LineData
}

type LineData struct {
	Data []string
}

func CreateXlsxSheet(_fielname string, _sheetList []SheetStruct) (bool, string) {
	if len(_sheetList) == 0 {
		return false, "Sheetlist cannot be empty."
	}
	//1. 创建文件钩子
	xlsx := excelize.NewFile()
	//2. 循环sheel列表
	_sheetLen := len(_sheetList)
	for s := 0; s < _sheetLen; s++ {
		if len(_sheetList[s].List) == 0 {
			return false, "Sheetlist cannot be empty"
		}
		//3. 创建sheet
		index := xlsx.NewSheet(_sheetList[s].SheetName)

		//4. 循环sheet中的行
		_lineLen := len(_sheetList[s].List)
		for l := 0; l < _lineLen; l++ {
			//5. 循环sheet中line的数据
			_data := _sheetList[s].List[l].Data
			_dataLen := len(_data)
			for d := 0; d < _dataLen; d++ {
				//6. 将数据导入到行中
				cellIndex := fmt.Sprintf("%s%d", I2A(d), l+1)
				if err := xlsx.SetCellValue(_sheetList[s].SheetName, cellIndex, _data[d]); err != nil {
					fmt.Println(err)
					return false, ""
				}
			}
		}
		//7. 将该sheet中所有数据添加完成后，创建sheet
		xlsx.SetActiveSheet(index)
	}

	//8. 将所有sheet写入文件中
	filePath := util.GetPath(_fielname + ".xlsx")
	if err := xlsx.SaveAs(filePath); err != nil {
		return false, err.Error()
	}
	return true, filePath
}
