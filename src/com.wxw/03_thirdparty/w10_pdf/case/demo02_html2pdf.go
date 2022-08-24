// @Time : 2022/8/24 18:01
// @Author : xiaoweiwei
// @File : demo02_html2pdf

package main

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"net/http"
	"strings"
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

	// Client code
	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	// The html string is also saved as base64 string in the JSON file
	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// The JSON can be saved, uploaded, etc.

	// Server code, create a new PDF generator from JSON, also looks for the wkhtmltopdf executable
	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Create the PDF
	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Use the PDF
	fmt.Printf("PDF size %d bytes", pdfgFromJSON.Buffer().Len())
}
