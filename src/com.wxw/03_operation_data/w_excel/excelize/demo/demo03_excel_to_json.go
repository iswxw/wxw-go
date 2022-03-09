/*
@Time : 2022/3/9 14:13
@Author : weixiaowei
@File : demo_excel_to_json
*/
package main

import (
	"encoding/json"
	"fmt"
	"framework/w_excel/excelize/common/util"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
)

func main() {
	xlsx, err := excelize.OpenFile(util.GetPath("cityfile.xlsx"))
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
	// Get all the rows in a sheet.
	rows, err := xlsx.GetRows("Sheet4")
	if err != nil {
		fmt.Println("err = ", err)
		os.Exit(1)
	}
	cities := make([]*City, 0)
	for _, row := range rows[1:] {
		city := CreateCity(row)
		cities = append(cities, &city)
	}

	byteCity, err := json.Marshal(cities)
	if err != nil {
		fmt.Println("err = ", err)
	}
	if err := ioutil.WriteFile(util.GetPath("a.txt"), byteCity, 0644); err != nil {
		return
	}
	//fmt.Printf("result = \n %s \n", string(byteCity))
}

func CreateCity(row []string) City {
	return City{
		CityId:        cast.ToInt64(row[0]),
		CityName:      row[1],
		CompanyCityId: cast.ToInt64(row[2]),
		CompanyId:     20001,
	}
}

// 定义结构体
type City struct {
	CityId        int64  `json:"city_id"`
	CityName      string `json:"city_name"`
	CompanyCityId int64  `json:"company_city_id"`
	CompanyId     int64  `json:"company_id"`
}
