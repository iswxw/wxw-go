// @Time : 2022/8/16 15:21
// @Author : xiaoweiwei
// @File : demo01_div

package main

import (
	"github.com/tiechui1994/gopdf"
	"github.com/tiechui1994/gopdf/core"
	"log"
)

// 测试文件来源 div_test.go
const (
	DivIg       = "IPAexG"
	DivMd       = "MPBOLD"
	DivStraight = 1 // 实线边框
	DIV_DASHED  = 2 // 虚线边框
	DivDotted   = 3 // 点状线的边框
	DivNone     = 4 // 无边框
)

func main() {

	rootPath := "src/com.wxw/03_thirdparty/w10_pdf/"

	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: DivIg,
		FileName: rootPath + "common/ttf/ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: DivMd,
		FileName: rootPath + "common/ttf/mplus-1p-bold.ttf",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2})
	r.SetPage("A4", "P")

	//r.RegisterExecutor(DivReportExecutor, core.Detail)
	r.RegisterExecutor(ComplexDivReportExecutor, core.Detail)
	r.Execute(rootPath + "tmp/div_test.pdf")
	//r.SaveAtomicCellText(rootPath + "tmp/div_test.txt")
}

func DivReportExecutor(report *core.Report) {
	font := core.Font{Family: DivMd, Size: 10}
	report.Font(DivMd, 10, "")
	report.SetFont(DivMd, 10)
	lineSpace := 1.0
	lineHeight := report.MeasureTextWidth("中")
	div := gopdf.NewDivWithWidth(300, lineHeight, lineSpace, report)

	div.SetFontWithColor(font, "11,123,244")
	//div.SetBackColor("11,123,244")
	div.RightAlign()
	div.SetMarign(core.NewScope(10, 20, 0, 0))
	div.SetBorder(core.NewScope(10, 50, 0, 0))
	div.SetContent(`Hello wxw-go`)
	div.GenerateAtomicCell()
}

func ComplexDivReportExecutor(report *core.Report) {
	font := core.Font{Family: DivMd, Size: 10}
	report.Font(DivMd, 10, "")
	report.SetFont(DivMd, 10)

	lineSpace := 1.0
	lineHeight := report.MeasureTextWidth("中")

	frame := gopdf.NewDivWithWidth(415, lineHeight, lineSpace, report)
	frame.SetFrameType(DivStraight)
	//frame.SetBackColor("222,111,11")
	frame.SetFont(font)
	frame.SetMarign(core.NewScope(10, 10, 0, 0))
	frame.SetBorder(core.NewScope(10, 10, 10, 10))
	content := `平台车险服务报告1 For a discussion of restrictions on subquery use, including performance issues for certain forms of subquery syntax, see Section C.4, “Restrictions on Subqueries”.`
	frame.SetContent(content)

	font2 := core.Font{Family: DivMd, Size: 16}
	frame2 := gopdf.NewDivWithWidth(415, lineHeight, lineSpace, report)
	frame2.SetFrameType(DivStraight)
	//frame.SetBackColor("222,111,11")
	frame2.SetFont(font2)
	frame2.SetMarign(core.NewScope(10, 10, 0, 0))
	frame2.SetBorder(core.NewScope(10, 10, 10, 10))
	content2 := `平台车险服务周报2`
	frame2.SetContent(content2)
	frame2.GenerateAtomicCell()
	if err := frame.GenerateAtomicCell(); err != nil {
		log.Println(err)
	}
}
