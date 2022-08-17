// @Time : 2022/8/17 10:59
// @Author : xiaoweiwei
// @File : quick_start

package main

import (
	"github.com/go-pdf/fpdf"
	"log"
)

// libs库：https://github.com/go-pdf/fpdf
// example: https://pkg.go.dev/github.com/go-pdf/fpdf#pkg-examples
func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
	}
}
