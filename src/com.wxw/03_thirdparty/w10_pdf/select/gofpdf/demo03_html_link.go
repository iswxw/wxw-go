// @Time : 2022/8/17 16:32
// @Author : xiaoweiwei
// @File : demo03_safety_secret

package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
)

func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	//var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"
	var imagePath = "src/com.wxw/03_thirdparty/w10_pdf/common/imgs/logo.png"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	//添加一页
	pdf.AddPage()
	//先设置字体
	pdf.SetFont("Helvetica", "", 20)
	_, lineHt := pdf.GetFontSize() //获取 line height
	fmt.Println(lineHt)            //7.055555555555556

	pdf.Write(lineHt, "To find out what's new in this tutorial, click ")
	pdf.SetFont("", "U", 0)
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)

	//Second page: image link and basic HTML with link
	pdf.AddPage()

	pdf.SetLink(link, 0, -1)
	pdf.Image(imagePath, 10, 12, 30, 0, false, "", 0, "www.google.com/ncr")
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(8)
	htmlStr := `You can now easily print text mixing different styles: <b>bold</b>, ` +
		`<i>italic</i>, <u>underlined</u>, or <b><i><u>all at once</u></i></b>!<br><br>` +
		`<center>You can also center text.</center>` +
		`<right>Or align it to the right.</right>` +
		`You can also insert links on text, such as ` +
		`<a href="http://www.fpdf.org">http://blog.csdn.net/wangshubo1989?viewmode=contents</a>, or on an image: click on the logo.`
	html := pdf.HTMLBasicNew()
	html.Write(5.0, htmlStr) //也可以手动指定 line height

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
		return
	}
}
