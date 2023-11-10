// @Time : 2022/8/17 19:04
// @Author : xiaoweiwei
// @File : demo10_book_sign.go

package main

import (
	"github.com/jung-kurt/gofpdf"
)

func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	//将字体加载进来
	//AddUTF8Font("给字体起个别名", "", "fontPath")
	pdf.AddUTF8Font("microsoft", "", fontPath)

	//使用这个字体
	//SetFont("字体的别名", "", size)
	pdf.SetFont("microsoft", "", 15)

	pdf.Bookmark("阿毛生信系统简介", 0, 0) //顶级书签，并且不显示该书签在 pdf 中
	pdf.Bookmark("发展历程", 1, -1)    //二级书签，名称叫“发展历程”，显示在 pdf 中
	pdf.Cell(0, 6, "发展历程")         //点击“发展历程”这个书签，即会跳转到 pdf 中“发展历程”的所在位置
	pdf.Ln(100)                    //空出几行，为了演示跳转效果
	pdf.Bookmark("荣誉资质", 1, -1)    //二级书签，名称叫“荣誉资质”，显示在 pdf 中
	pdf.Cell(0, 6, "荣誉资质")         //点击“荣誉资质”这个书签，即会跳转到 pdf 中“荣誉资质”的所在位置

	pdf.AddPage()
	pdf.Bookmark("我们的产品", 0, 0)
	pdf.Bookmark("核心产品简介", 1, -1)
	pdf.Cell(0, 6, "核心产品简介")

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		return
	}
}
