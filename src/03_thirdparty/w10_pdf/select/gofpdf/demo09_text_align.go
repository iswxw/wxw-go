// @Time : 2022/8/17 18:59
// @Author : xiaoweiwei
// @File : demo09_text_align

package main

import (
	"github.com/jung-kurt/gofpdf"
	"strings"
)

// 文本对齐
func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.SetLeftMargin(50.0)
	pdf.SetRightMargin(50.0)
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 12)

	pdf.WriteAligned(0, 35, "This text is the default alignment, Left", "")
	pdf.Ln(35)

	pdf.WriteAligned(0, 35, "This text is aligned Left", "L")
	pdf.Ln(35)

	pdf.WriteAligned(0, 35, "This text is aligned Center", "C")
	pdf.Ln(35)

	line := "This can by used to write justified text"
	leftMargin, _, rightMargin, _ := pdf.GetMargins()
	pageWidth, _ := pdf.GetPageSize()
	pageWidth -= leftMargin + rightMargin

	pdf.SetWordSpacing((pageWidth - pdf.GetStringWidth(line)) / float64(strings.Count(line, " ")))
	pdf.WriteAligned(pageWidth, 35, line, "L")

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		return
	}
}
