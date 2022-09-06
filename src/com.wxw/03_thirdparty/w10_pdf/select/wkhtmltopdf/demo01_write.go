// @Time : 2022/8/24 15:20
// @Author : xiaoweiwei
// @File : main

package main

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// https://pkg.go.dev/github.com/andrewcharlton/wkhtmltopdf-go
func main() {

	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	var fontPath = "src/com.wxw/03_thirdparty/w10_pdf/common/ttf/microsoft.ttf"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	//添加一页
	pdf.AddPage()

	//写文字内容之前，必须先要设置好字体
	pdf.SetFont("Arial", "B", 16)

	//CellFormat: 表格显示样式设置
	//CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(0, 0, "Welcome to golang code.com", "0", 0, "LM", false, 0, "")

	//将字体加载进来
	//AddUTF8Font("给字体起个别名", "", "fontPath")
	pdf.AddUTF8Font("microsoft", "", fontPath)

	//使用这个字体
	//SetFont("字体的别名", "", size)
	pdf.SetFont("microsoft", "", 20)
	htmlBasicNew := pdf.HTMLBasicNew()
	fontSize, _ := pdf.GetFontSize()
	htmlBasicNew.Write(fontSize, getTagHTML(rootPath))

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		return
	}
}

func ParseHtml(path string) bytes.Buffer {
	t, err := template.ParseFiles(path + "common/tmpl/test.html")
	if err != nil {
		log.Printf("Parse template failed, err%v\n", err)
		return bytes.Buffer{}
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, nil); err != nil {
		log.Printf("err%v\n", err)
		return bytes.Buffer{}
	}
	return buf
}

func getTagHTML(path string) string {

	file, err := os.Open(path + "common/tmpl/test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}
