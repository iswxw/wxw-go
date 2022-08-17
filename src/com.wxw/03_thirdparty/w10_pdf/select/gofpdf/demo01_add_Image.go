// @Time : 2022/8/17 15:46
// @Author : xiaoweiwei
// @File : demo01_add_Image

package main

import "github.com/jung-kurt/gofpdf"

// PDF 添加图片
func main() {

	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"

	//设置页面参数
	pdf := gofpdf.New("P", "mm", "A4", "")

	//添加一页
	pdf.AddPage()

	//将图片放入到 pdf 文档中
	//ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		rootPath+"/common/imgs/gopher.jpg", 0, 0, 0, 0, false,
		gofpdf.ImageOptions{ImageType: "jpg", ReadDpi: false}, 0, "",
	)

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		return
	}

}
