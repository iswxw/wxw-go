// @Time : 2022/8/24 15:20
// @Author : xiaoweiwei
// @File : main

package main

import (
	"fmt"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// https://github.com/SebastiaanKlippert/go-wkhtmltopdf/issues/28
func main() {
	rootPath := "src/com.wxw/03_thirdparty/w10_pdf/"
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {

		return
	}
	page := wkhtml.NewPageReader(strings.NewReader(getTagHTMLMain()))
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("workingDir = ", workingDir)
	page.Allow.Set(workingDir)

	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	if err = pdfg.WriteFile(rootPath + "tmp/hello.pdf"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")

}

func getTagHTMLMain() string {
	RootPath := "src/com.wxw/03_thirdparty/w10_pdf/"
	file, err := os.Open(RootPath + "common/tmpl/test.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	return string(b)
}
