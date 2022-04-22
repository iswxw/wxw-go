/*
@Time : 2022/4/22 20:13
@Author : weixiaowei
@File : demo05_gen_sql
*/
package main

import (
	"fmt"
	"framework/w_excel/excelize/common/util"
	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx2, _ := excelize.OpenFile(util.GetPath("test1.xlsx"))
	rows1, _ := xlsx2.GetRows("Sheet1")

	updateBillSql := "update bill_%d set create_time = %s where company_id = 20002 and item_id = 1100"
	for _, a := range rows1 {
		fmt.Println(fmt.Sprintf(updateBillSql, a[0]))
	}
}
