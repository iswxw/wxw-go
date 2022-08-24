/*
@Time: 2022/8/14 22:56
@Author: wxw
@File: demo01_generate_pdf
*/
package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	rootPath := "src/com.wxw/03_thirdparty/w10_pdf/"
	err := pdf.AddTTFFont("wts11", rootPath+"common/ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetTextColor(156, 197, 140)
	pdf.Cell(nil, "您好")

	pdf.WritePdf(rootPath + "tmp/hello.pdf")
}
