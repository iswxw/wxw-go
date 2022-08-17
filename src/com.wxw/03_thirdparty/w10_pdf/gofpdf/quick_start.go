// @Time : 2022/8/17 15:36
// @Author : xiaoweiwei
// @File : quick_start

package main

import "github.com/jung-kurt/gofpdf"

// 相关资料：https://blog.csdn.net/zhangyibei2008/article/details/107057883
func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	//添加一页
	pdf.AddPage()

	//写文字内容之前，必须先要设置好字体
	pdf.SetFont("Arial", "B", 16)

	//CellFormat: 表格显示样式设置
	//CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(0, 0, "Welcome to golang code.com", "0", 0, "LM", false, 0, "")

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		return
	}
}
