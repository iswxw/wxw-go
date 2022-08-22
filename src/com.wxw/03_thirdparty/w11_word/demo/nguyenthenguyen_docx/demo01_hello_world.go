/*
@Time: 2022/8/16 22:46
@Author: wxw
@File: demo01_hello_world
*/
package main

import (
	"github.com/nguyenthenguyen/docx"
	"strconv"
)

func main() {
	// Read from docx file
	r, err := docx.ReadDocxFile("./template.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)

	// Or read from a filesystem object:
	// r, err := docx.ReadDocxFromFS(file string, fs fs.FS)

	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	docx1.Replace("old_1_1", "new_1_1", -1)
	docx1.Replace("old_1_2", "new_1_2", -1)
	docx1.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx", 1)
	docx1.ReplaceHeader("out with the old", "in with the new")
	docx1.ReplaceFooter("Change This Footer", "new footer")
	docx1.WriteToFile("./new_result_1.docx")

	docx2 := r.Editable()
	docx2.Replace("old_2_1", "new_2_1", -1)
	docx2.Replace("old_2_2", "new_2_2", -1)
	docx2.WriteToFile("./new_result_2.docx")

	// Or write to ioWriter
	// docx2.Write(ioWriter io.Writer)

	docx3 := r.Editable()
	//Currently only swaps apples for apples i.e. png to png, and not png to jpeg etc.
	d.ReplaceImage("word/media/image1.png", "./new.png")
	docx3.WriteToFile("./new_result_3.docx")

	// replace the last image
	imageIndex := docxRaw.ImagesLen()
	d.ReplaceImage("word/media/image"
	strconv.Itoa(imageIndex) + ".png", "./new.png")

	r.Close()
}
