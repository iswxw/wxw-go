// @Time : 2022/8/17 10:59
// @Author : xiaoweiwei
// @File : quick_start

package main

import (
	"github.com/go-pdf/fpdf"
	"log"
)

// libs库：https://github.com/go-pdf/fpdf
// example: https://pkg.go.dev/github.com/go-pdf/fpdf#pkg-examples
func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	//将字体加载进来
	//AddUTF8Font("给字体起个别名", "", "fontPath")
	pdf.AddUTF8Font("microsoft", "", fontPath)

	//使用这个字体
	//SetFont("字体的别名", "", size)
	pdf.SetFont("microsoft", "", 20)
	pdf.Text(5, 10, "横看成岭侧成峰")
	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
	}
}
