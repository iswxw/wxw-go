// @Time : 2022/8/24 18:01
// @Author : xiaoweiwei
// @File : demo02_html2pdf

package main

import (
	"bytes"
	"html/template"
	//"github.com/pleximus/go-htmltopdf"
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
	const html = `<!doctype html><html><head><title>WKHTMLTOPDF TEST</title></head><body>HELLO PDF</body></html>`

}
