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

	basePath := "D:\\Project\\wxw-go\\src\\com.wxw\\03_thirdparty\\w10_pdf\\"

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", basePath+"common\\ttf\\wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf(basePath + "tmp\\hello.pdf")

}
