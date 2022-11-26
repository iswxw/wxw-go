// @Time : 2022/8/17 17:23
// @Author : xiaoweiwei
// @File : demo7_multi_column

package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

func main() {

	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"
	//var imagePath = "src/com.wxw/03_thirdparty/w10_pdf/common/imgs/logo.png"

	var y0 float64
	var crrntCol int

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetDisplayMode("fullpage", "TwoColumnLeft")

	//将字体加载进来
	//AddUTF8Font("给字体起个别名", "", "fontPath")
	pdf.AddUTF8Font("microsoft", "", fontPath)

	// 设置作者
	pdf.SetAuthor("小伟", true)

	// 设置标题
	titleStr := "技术能量站"
	pdf.SetTitle(titleStr, true)

	// 给指定列设置位置
	setCol := func(col int) {
		//根据给定的列，设置位置
		crrntCol = col
		x := 10.0 + float64(col)*65.0
		pdf.SetLeftMargin(x)
		pdf.SetX(x)
	}

	// 设置标题
	chapterTitle := func(chapNum int, titleStr string) {
		pdf.SetFont("microsoft", "", 12)
		pdf.SetFillColor(200, 220, 255) //background color
		pdf.CellFormat(0, 6, fmt.Sprintf("%s", titleStr),
			"", 1, "L", true, 0, "",
		)
		pdf.Ln(2)
		y0 = pdf.GetY()
	}

	// 章节内容
	chapterBody := func(fileStr string) {
		txtBuf, err := ioutil.ReadFile(fileStr)
		if err != nil {
			panic(err)
		}
		pdf.MultiCell(60, 5, string(txtBuf), "", "", false)
		pdf.Cell(0, 5, "(end of excerpt)")

		setCol(0) //返回第一列
	}

	// 打印章节信息
	printChapter := func(num int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(num, titleStr)
		chapterBody(fileStr)
	}

	// 判断是否需要分页
	pdf.SetAcceptPageBreakFunc(func() bool {
		if crrntCol < 2 {
			setCol(crrntCol + 1)
			pdf.SetY(y0)
			return false //继续保持在当前页
		}
		setCol(0)

		return true //执行分页
	})

	// 设置页眉
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("microsoft", "", 15) //设置“仿宋”字体
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		pdf.SetDrawColor(0, 80, 180)  //frame color
		pdf.SetFillColor(230, 230, 0) //background color
		pdf.SetTextColor(220, 50, 50) //text color
		pdf.SetLineWidth(1)
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
		pdf.Ln(5)
		y0 = pdf.GetY() //保存纵坐标
	})

	// 设置页脚
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(
			0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "",
		)
	})

	printChapter(1, "第一章，项目背景", rootPath+"common/text/20k_c2.txt")
	printChapter(2, "第二章，项目背景", rootPath+"common/text/20k_c1.txt")
	//printChapter(2, "A RUNAWAY REEF", rootPath+"common/text/20k_c1.txt")

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
		return
	}

}
