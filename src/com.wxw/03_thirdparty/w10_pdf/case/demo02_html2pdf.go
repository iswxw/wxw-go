// @Time : 2022/8/24 18:01
// @Author : xiaoweiwei
// @File : demo02_html2pdf

package main

import (
	"bytes"
	"github.com/jung-kurt/gofpdf"
	"html/template"
	"net/http"
)

const page = `
<html>
  <body>
    <h1>Test Page</h1>

	<p>Path: {{.}}</p>
  </body>
</html>`

func main() {
	GeneratorPdf()
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("page").Parse(page))
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, r.URL.String())

	//w.Header().Set("Content-Type", "application/pdf")
	//w.Header().Set("Content-Disposition", `attachment; filename="test.pdf"`)
}

// 需要安装wkhtmltopdf环境
func GeneratorPdf() {

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

	htmlBasicNew := pdf.HTMLBasicNew()
	fontSize, _ := pdf.GetFontSize()
	htmlBasicNew.Write(fontSize, page)

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		return
	}
}
