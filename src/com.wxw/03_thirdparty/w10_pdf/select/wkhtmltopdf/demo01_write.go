// @Time : 2022/8/24 15:20
// @Author : xiaoweiwei
// @File : main

package main

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// https://pkg.go.dev/github.com/andrewcharlton/wkhtmltopdf-go

var RootPath = "src/com.wxw/03_thirdparty/w10_pdf/"

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done

}

func ParseHtml() bytes.Buffer {
	t, err := template.ParseFiles(RootPath + "common/tmpl/test.html")
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

func GeneratorPdf() {

}

func getTagHTML() string {
	file, err := os.Open(RootPath + "common/tmpl/test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}
