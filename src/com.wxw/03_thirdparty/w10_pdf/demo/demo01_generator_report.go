// @Time : 2022/8/16 14:38
// @Author : xiaoweiwei
// @File : demo01_generator_report

package main

import (
	"fmt"
	"github.com/tiechui1994/gopdf/core"
)

const (
	TableIg = "IPAexG"
	TableMd = "MPBOLD"
	TableMy = "微软雅黑"
)

var (
	basePath = "src/com.wxw/03_thirdparty/w10_pdf/"
)

func main() {

	report := core.CreateReport()

	// 设置字体
	font1 := core.FontMap{
		FontName: TableMy,
		FileName: basePath + "common/ttf/microsoft.ttf",
	}
	report.SetFonts([]*core.FontMap{&font1})

	report.SetPage("A4", "P")

	report.RegisterExecutor(SimpleTableExecutor, core.Detail)

	report.Execute(basePath + "tmp/simple_table.pdf")
	fmt.Println("当前页码 = ", report.GetCurrentPageNo())
}

func SimpleTableExecutor(report *core.Report) {

	//lineSpace := 1.0
	//lineHeight := report.MeasureTextWidth("中")

}
