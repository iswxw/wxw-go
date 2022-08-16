// @Time : 2022/8/16 16:21
// @Author : xiaoweiwei
// @File : demo03_links

package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("times", rootPath+"./common/ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("times", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetXY(30, 40)
	pdf.SetTextColor(0, 255, 0)
	pdf.Text("Link to example.com")
	pdf.AddExternalLink("http://example.com/", 27.5, 28, 125, 15)

	pdf.SetXY(30, 70)
	pdf.Text("Link to second page")
	pdf.AddInternalLink("anchor", 27.5, 58, 120, 15)

	pdf.AddPage()
	pdf.SetXY(30, 100)
	pdf.SetAnchor("anchor")
	pdf.Text("Anchor position")

	pdf.WritePdf(rootPath + "tmp/hello.tmp.pdf")

}
