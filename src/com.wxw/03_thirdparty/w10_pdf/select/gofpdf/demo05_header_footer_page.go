// @Time : 2022/8/17 16:41
// @Author : xiaoweiwei
// @File : demo05_header_footer

package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

//  主要演示 页眉、页脚和分页
func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"
	var imagePath = "src/com.wxw/03_thirdparty/w10_pdf/common/imgs/logo.png"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	//将字体加载进来
	//AddUTF8Font("给字体起个别名", "", "fontPath")
	pdf.AddUTF8Font("microsoft", "", fontPath)
	//使用这个字体
	//SetFont("字体的别名", "", size)
	pdf.SetFont("microsoft", "", 20)

	//设置页眉
	pdf.SetHeaderFuncMode(func() {
		pdf.Image(imagePath, 0, 0, 0, 0, false, "", 0, "")
		pdf.SetY(5)
		pdf.Ln(10)
	}, true)

	//设置页脚
	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.CellFormat(
			0, 10,
			fmt.Sprintf("当前第 %d 页，共 {nb} 页", pdf.PageNo()), //字符串中的 {nb}。大括号是可以省的，但不建议这么做
			"", 0, "C", false, 0, "",
		)
	})

	//给个空字符串就会去替换默认的 "{nb}"。
	//如果这里指定了特别的字符串，那么SetFooterFunc() 中的 "nb" 也必须换成这个特别的字符串
	pdf.AliasNbPages("")

	pdf.AddPage()
	for j := 0; j < 100; j++ {
		pdf.CellFormat(0, 10, fmt.Sprintf("正在打印：%d", j),
			"", 1, "", false, 0, "",
		)
	}

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
		return
	}
}
