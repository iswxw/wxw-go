// @Time : 2022/8/24 15:11
// @Author : xiaoweiwei
// @File : demo02_byte_array

package main

import (
	"bytes"
	"github.com/go-pdf/fpdf"
	"html/template"
	"log"
)

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
	pdf.SetFont("microsoft", "", 10)
	//pdf.Text(5, 10, "横看成岭侧成峰")

	_, lineHt := pdf.GetFontSize()
	html := ParseHtml()
	basicNew := pdf.HTMLBasicNew()
	basicNew.Write(lineHt, html.String())
	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
	}
}

func ParseHtml() bytes.Buffer {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	t, err := template.ParseFiles(rootPath + "common/tmpl/test.html")
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
