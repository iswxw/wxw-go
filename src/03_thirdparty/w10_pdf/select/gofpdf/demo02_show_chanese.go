// @Time : 2022/8/17 16:15
// @Author : xiaoweiwei
// @File : demo02_show_chanese

package main

import (
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")
	//添加一页
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
		return
	}

}
