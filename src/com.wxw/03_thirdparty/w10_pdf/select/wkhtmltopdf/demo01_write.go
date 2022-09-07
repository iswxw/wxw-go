// @Time : 2022/8/24 15:20
// @Author : xiaoweiwei
// @File : main

package main

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// https://pkg.go.dev/github.com/andrewcharlton/wkhtmltopdf-go
func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	url := "http://www.baidu.com/"
	pdf.AddPage(wkhtmltopdf.NewPage(url))

	//html := ParseHtml(rootPath)
	//pdf.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(html.Bytes())))

	if err := pdf.Create(); err != nil {
		return
	}
	if err := pdf.WriteFile(rootPath + "tmp/hello.pdf"); err != nil {
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
