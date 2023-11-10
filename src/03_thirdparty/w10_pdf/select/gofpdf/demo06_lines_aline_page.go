// @Time : 2022/8/17 16:50
// @Author : xiaoweiwei
// @File : demo06_lines_aline_page

package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
)

// 主要演示 多行文字段落、行对齐、分页
func main() {

	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"
	//var imagePath = "src/com.wxw/03_thirdparty/w10_pdf/common/imgs/logo.png"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")
	w, h := pdf.GetPageSize()
	fmt.Printf("pdf size, w:%.2f, h:%.2f \n", w, h) //pdf size, w:210.00, h:297.00

	titleStr := "车险服务周报"

	//将字体加载进来
	//AddUTF8Font("给字体起个别名", "", "fontPath")
	pdf.AddUTF8Font("microsoft", "", fontPath)

	//使用这个字体
	//SetFont("字体的别名", "", size)
	pdf.SetFont("microsoft", "", 14)

	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)

	//设置页眉
	pdf.SetHeaderFuncMode(func() {
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetY(0.6)            //先要设置 Y，然后再设置 X。否则，会导致 X 失效
		pdf.SetX((210 - wd) / 2) //水平居中的算法

		pdf.SetDrawColor(0, 80, 180)  //frame color
		pdf.SetFillColor(230, 230, 0) //background color
		pdf.SetTextColor(220, 50, 50) //text color

		pdf.SetLineWidth(1)

		pdf.CellFormat(wd, 10, titleStr, "1", 1, "CM", true, 0, "")

		// 第 5 个参数，实际效果是：指定下一行的位置
		// Ln(h float64) 表示：创建一个高度为 h 的空行。
		pdf.Ln(5)
		// 0：表示不换行，并紧跟在这个 Cell 的右边。
		// 1：发生换行，并在下一行的顶头位置。
		// 2：发生换行，但是会在这个 Cell 的下方。

	}, false)

	//设置页脚
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetTextColor(128, 128, 128)
		pdf.CellFormat(
			0, 5,
			fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "",
		)
	})

	// 设置标题
	// chapNum 章节数据
	// titleStr 章节标题
	chapterTitle := func(chapNum int, titleStr string) {
		pdf.SetFillColor(200, 220, 255) //background color
		pdf.CellFormat(0, 6, fmt.Sprintf("%s", titleStr), "", 1, "L", true, 0, "")

		pdf.Ln(2)
	}

	// 设置主体
	chapterBody := func(fileStr string) {
		textStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}

		//输出对齐文本
		pdf.MultiCell(0, 5, string(textStr), "", "", false)
		pdf.Ln(-1)
		//pdf.SetFont("microsoft", "I", 0)
		pdf.Cell(0, 5, "本章节加载完毕")
	}

	//印刷每一页
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)
		chapterBody(fileStr)
	}

	printChapter(1, "第一章，项目背景", rootPath+"common/text/20k_c2.txt")
	//printChapter(2, "A RUNAWAY REEF", rootPath+"common/text/20k_c1.txt")

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
		return
	}

}
