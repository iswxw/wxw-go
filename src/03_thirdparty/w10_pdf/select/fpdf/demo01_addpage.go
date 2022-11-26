// @Time : 2022/8/17 11:45
// @Author : xiaoweiwei
// @File : demo01_addpage

package main

import (
	"fmt"
	"github.com/go-pdf/fpdf"
)

func main() {

	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetTopMargin(30)
	pdf.SetHeaderFuncMode(func() {
		pdf.Image(fmt.Sprintf("%scommon/imgs/logo.png", rootPath), 10, 6, 30, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 15)
		pdf.Cell(80, 0, "")
		pdf.CellFormat(30, 10, "Title", "1", 0, "C", false, 0, "")
		pdf.Ln(20)
	}, true)
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Times", "", 12)
	for j := 1; j <= 40; j++ {
		pdf.CellFormat(0, 10, fmt.Sprintf("Printing line number %d", j),
			"", 1, "", false, 0, "")
	}
	if err := pdf.OutputFileAndClose(rootPath + "/tmp/Fpdf_AddPage.pdf"); err != nil {
		return
	}
}
