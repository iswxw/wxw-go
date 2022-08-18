/*
@Time : 2022/3/7 18:48
@Author : weixiaowei
@File : demo_json_to_excel
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"src/com.wxw/project_actual/src/com.wxw/03_thirdparty/w05_excel/common/util"
	"strconv"
)

func main() {
	jsonStr := `{"1":{"city_id":1,"city_name_cn":"北京市"}}`
	cityMaps := make(map[int32]*CityMap, 300)
	if err := json.Unmarshal([]byte(jsonStr), &cityMaps); err != nil {
		_ = fmt.Errorf("err = %s \n", err)
	}
	header := []interface{}{"城市编号", "城市名称"}

	data := makeExcelMonthData(cityMaps)
	// fmt.Println("jsonStr = ", jsonStr)
	ExcelExportDataT0Excel(header, data, "Sheet1")
}

func ExcelExportDataT0Excel(header []interface{}, data [][]interface{}, sheetName string) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet(sheetName)

	if err := f.SetSheetRow(sheetName, "F1", &header); err != nil {
		log.Println("err1= ", err)
		return
	}
	axi := 1
	for _, obj := range data {
		axi = axi + 1
		if err := f.SetSheetRow(sheetName, "F"+strconv.Itoa(axi), &obj); err != nil {
			log.Println("err2= ", err)
			return
		}
	}
	f.SetActiveSheet(index)
	//timeStr := time.Now().Format("2006-01-02 15:04:05")
	// Save xlsx file by the given path.
	if err := f.SaveAs(util.GetPath("城市代码表-市级.xlsx")); err != nil {
		fmt.Println(err)
	}
}

// 构造月对账数据
func makeExcelMonthData(data map[int32]*CityMap) [][]interface{} {
	var results [][]interface{}
	for _, value := range data {
		result := ConvertStructToStr(value)
		results = append(results, result)
	}
	return results
}

func ConvertStructToStr(value *CityMap) []interface{} {
	results := make([]interface{}, 4)
	results[0] = value.CityId
	results[1] = value.CityNameCn
	return results
}

type CityMap struct {
	CityId     int    `json:"city_id"`
	CityNameCn string `json:"city_name_cn"`
}
