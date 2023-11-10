// @Time : 2022/8/16 16:02
// @Author : xiaoweiwei
// @File : demo02_

package main

import (
	"github.com/signintech/gopdf"
	"log"
)

func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("loma", rootPath+"common/ttf/Loma.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Image(rootPath+"common/imgs/gopher.jpg", 200, 50, nil) //print image
	err = pdf.SetFont("loma", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetXY(250, 200)                //move current location
	pdf.Cell(nil, "gopher and gopher") //print text

	pdf.WritePdf(rootPath + "tmp/image.pdf")
}
