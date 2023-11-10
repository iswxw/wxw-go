// @Time : 2022/8/17 17:41
// @Author : xiaoweiwei
// @File : demo08_table

package main

import (
	"bufio"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"os"
	"strings"
)

func main() {
	var rootPath = "src/com.wxw/03_thirdparty/w10_pdf/"
	pdf := gofpdf.New("P", "mm", "A4", "") // 210 x 297
	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}

	countryList := make([]countryType, 0, 8)
	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}

	loadData := func(fileStr string) {
		fl, err := os.Open(fileStr)
		if err != nil {
			log.Println(err)
			panic(err)
		}

		defer fl.Close()

		scanner := bufio.NewScanner(fl)
		var c countryType
		for scanner.Scan() {
			//Austria;Vienna;83859;8075
			lineStr := scanner.Text()
			list := strings.Split(lineStr, ";")

			if len(list) == 4 {
				c.nameStr = list[0]
				c.capitalStr = list[1]
				c.areaStr = list[2]
				c.popStr = list[3]
				countryList = append(countryList, c)
			} else {
				fmt.Printf("error tokenizing %s\n", lineStr)
			}
		}

		if len(countryList) == 0 {
			fmt.Printf("error loading data from %s\n", fileStr)
		}
	}

	//简单的表格
	basicTable := func() {
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)

		for _, str := range header {
			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
		}

		pdf.Ln(-1)

		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}

	//Better table
	improvedTable := func() {
		w := []float64{40.0, 35.0, 40.0, 45.0} //列宽
		wSum := 0.0

		for _, v := range w {
			wSum += v
		}

		left := (210 - wSum) / 2
		pdf.SetX(left)

		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
		}

		pdf.Ln(-1)

		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3), "LR", 0, "", false, 0, "")
			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3), "LR", 0, "", false, 0, "")
			pdf.Ln(-1)
		}

		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	//Colored table
	fancyTable := func() {
		pdf.SetFillColor(255, 0, 0)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(128, 0, 0)
		pdf.SetLineWidth(.3)
		pdf.SetFont("", "B", 0)

		w := []float64{40, 35, 40, 45}
		wSum := 0.0

		for _, v := range w {
			wSum += v
		}

		left := (210 - wSum) / 2
		pdf.SetX(left)

		for j, str := range header {
			pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
		}

		pdf.Ln(-1)

		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetFont("", "", 0)

		fill := false
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3), "LR", 0, "", fill, 0, "")
			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3), "LR", 0, "", fill, 0, "")
			pdf.Ln(-1)
			fill = !fill
		}
		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	}

	loadData(rootPath + "common/text/countries.txt")
	pdf.SetFont("Arial", "", 14)

	pdf.AddPage()
	basicTable()

	pdf.AddPage()
	improvedTable()

	pdf.AddPage()
	fancyTable()

	if err := pdf.OutputFileAndClose(rootPath + "tmp/hello.pdf"); err != nil {
		log.Println(err)
		return
	}
}

//strDelimit converts 'ABCDEFG' to, for example, 'A,BCD,EFG'
func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}
