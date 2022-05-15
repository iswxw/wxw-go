/*
@Time : 2022/1/23 23:31
@Author : weixiaowei
@File : demo01_create_xlsx
@link: https://xuri.me/excelize/zh-hans/base/installation.html#NewFile
*/
package main

import (
	"fmt"
	"framework/w05_excel/excelize/common/util"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

func main() {
	// 1. 创建 使用 NewFile 新建 Excel 工作薄，新创建的工作簿中会默认包含一个名为 Sheet1 的工作表
	f := excelize.NewFile()

	// 2. 设置工作表名称
	//根据给定的新旧工作表名称（大小写敏感）重命名工作表。工作表名称最多允许使用 31 个字符，
	//此功能仅更改工作表的名称，而不会更新与单元格关联的公式或引用中的工作表名称。
	//因此使用此功能重命名工作表后可能导致公式错误或参考引用问题。
	sheetName := "学生成绩单"
	f.SetSheetName("Sheet1", sheetName) //设置工作表的名称

	// 3. 准备数据
	grade := [][]interface{}{
		{"考试成绩统计表"},
		{"考试名称:期中考试", nil, nil, "文综", nil, nil, "理综"},
		{"序号", "学号", "姓名", "历史", "地理", "政治", "生物", "化学", "物理", "总分", "平均分"},
		{1, "1001", "青雉1", 11, 22, 33, 44, 55, 66, nil, nil},
		{2, "1002", "青雉2", 11, 22, 33, 44, 55, 66, nil, nil},
		{3, "1003", "青雉3", 11, 22, 33, 44, 55, 66, nil, nil},
		{4, "1004", "青雉4", 11, 22, 33, 44, 55, 66, nil, nil},
		{5, "1005", "青雉5", 11, 22, 33, 44, 55, 66, nil, nil},
		{6, "1006", "青雉6", 11, 22, 33, 44, 55, 66, nil, nil},
	}

	for i, obj := range grade {

		// 4.根据行和列拼接单元格名称
		name, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Println(fmt.Sprintf("拼接单元格名称失败,错误:%s", err))
			return
		}

		// 5. 按行赋值
		//根据给定的工作表名称（大小写敏感）、起始坐标和 slice 类型引用按行赋值。
		//例如，在名为 Sheet1 的工作簿第 6 行上，以 B6 单元格作为起始坐标按行赋值：
		//err := f.SetSheetRow("Sheet1", "B6", &[]interface{}{"1", nil, 2})
		err = f.SetSheetRow(sheetName, name, &obj)
		if err != nil {
			fmt.Println(fmt.Sprintf("按行写入数据失败,错误:%s", err))
			return
		}
	}

	// 6. 设置公式
	//根据给定的工作表名（大小写敏感）和单元格坐标设置该单元格上的公式。
	//公式的结果可在工作表被 Office Excel 应用程序打开时计算，
	//或通过 CalcCellValue 函数计算单元格的值。
	//若 Excel 应用程序打开工作簿后未对设置的单元格公式进行计算，请在设置公式后调用 UpdateLinkedValue 清除单元格缓存。
	ref := "J4:J9"
	shared := excelize.STCellFormulaTypeShared
	formulaOpts := excelize.FormulaOpts{Type: &shared, Ref: &ref}
	err := f.SetCellFormula(sheetName, "J4", "=SUM(D4:I4)", formulaOpts)
	if err != nil {
		fmt.Println(fmt.Sprintf("设置公式失败,错误:%s", err))
		return
	}

	// 7. 清除单元格缓存
	err = f.UpdateLinkedValue()
	if err != nil {
		fmt.Println(fmt.Sprintf("清除单元格缓存失败,错误:%s", err))
		return
	}

	//设置相同的公式后发现没有应用上,因此每个都添加公式,求大佬解释,为啥我上面的代码设置相同的公式失败了
	//起初原以为是修改了工作表名称,后面测试并不是,难道是WPS的问题?疑惑！
	for i := 5; i <= 9; i++ {
		index := strconv.Itoa(i)
		err = f.SetCellFormula(sheetName, "J"+index, fmt.Sprintf("=SUM(D%s:I%s)", index, index), formulaOpts)
		if err != nil {
			fmt.Println(fmt.Sprintf("设置公式失败,错误:%s", err))
			return
		}
	}

	// 8. 合并单元格
	//根据给定的工作表名（大小写敏感）和单元格坐标区域合并单元格。合并区域内仅保留左上角单元格的值，其他单元格的值将被忽略。
	//例如，合并名为 Sheet1 的工作表上 D3:E9 区域内的单元格：
	//err := f.MergeCell("Sheet1", "D3", "E9")
	//如果给定的单元格坐标区域与已有的其他合并单元格相重叠，已有的合并单元格将会被删除。
	err = f.MergeCell(sheetName, "A1", "K1")
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}

	err = f.MergeCell(sheetName, "A2", "C2")
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}

	err = f.MergeCell(sheetName, "D2", "F2")
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}

	err = f.MergeCell(sheetName, "G2", "I2")
	if err != nil {
		fmt.Println(fmt.Sprintf("合并单元格失败,错误:%s", err))
		return
	}

	//--设置单元格样式
	//func (f *File) SetCellStyle(sheet, hcell, vcell string, styleID int) error
	//根据给定的工作表名、单元格坐标区域和样式索引设置单元格的值
	//。样式索引可以通过 NewStyle 函数获取。
	//注意，在同一个坐标区域内的 diagonalDown 和 diagonalUp 需要保持颜色一致。
	//SetCellStyle 将覆盖单元格的已有样式，而不会将样式与已有样式叠加或合并。
	styleCenter, err := f.NewStyle(&excelize.Style{
		Border: nil,
		Fill:   excelize.Fill{},
		Font:   nil,
		Alignment: &excelize.Alignment{
			Horizontal:      "center", //水平居中
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "", //垂直居中
			WrapText:        false,
		},
		Protection:    nil,
		NumFmt:        0,
		DecimalPlaces: 0,
		CustomNumFmt:  nil,
		Lang:          "",
		NegRed:        false,
	})

	style2, err := f.NewStyle(&excelize.Style{
		Border: nil,
		Fill: excelize.Fill{
			Type:    "pattern", //纯色填充
			Pattern: 1,
			Color:   []string{"DFEBF6"},
			Shading: 0,
		},
		Font: nil,
		Alignment: &excelize.Alignment{
			Horizontal:      "center", //水平居中
			Indent:          0,
			JustifyLastLine: false,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "", //垂直居中
			WrapText:        false,
		},
		Protection:    nil,
		NumFmt:        0,
		DecimalPlaces: 0,
		CustomNumFmt:  nil,
		Lang:          "",
		NegRed:        false,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("设置样式失败,错误:%s", err))
		return
	}
	_ = f.SetCellStyle(sheetName, "A1", "A1", style2)
	_ = f.SetCellStyle(sheetName, "A2", "A2", styleCenter)
	_ = f.SetCellStyle(sheetName, "D2", "D2", styleCenter)
	_ = f.SetCellStyle(sheetName, "G2", "G2", styleCenter)

	//设置列宽度
	//func (f *File) SetColWidth(sheet, startcol, endcol string, width float64) error
	//根据给定的工作表名称（大小写敏感）、列范围和宽度值设置单个或多个列的宽度。例如设置名为 Sheet1 工作表上 A 到 H 列的宽度为 20：
	//f := excelize.NewFile()
	//err := f.SetColWidth("Sheet1", "A", "H", 20)
	_ = f.SetColWidth(sheetName, "D", "D", 7)

	//--创建表格
	//根据给定的工作表名、单元格坐标区域和条件格式创建表格。
	//注意，表格坐标区域至少需要包含两行：字符型的标题行和内容行。
	//每列标题行的字符需保证是唯一的，并且必须在调用 AddTable 函数前设置表格的标题行数据。多个表格的坐标区域不能有交集。
	//可选参数 table_name 用以设置自定义表格名称，同一个工作表内的表格名称应该是唯一的。
	//Excelize 支持的表格样式 table_style 参数：
	//TableStyleLight1 - TableStyleLight21
	//TableStyleMedium1 - TableStyleMedium28
	//TableStyleDark1 - TableStyleDark11
	err = f.AddTable(sheetName, "A3", "K9", `{"table_name":"表格","table_style":"TableStyleLight21"}`)
	if err != nil {
		fmt.Println(fmt.Sprintf("设置样式失败,错误:%s", err))
		return
	}

	//--另存为
	//使用 SaveAs 保存 Excel 文档为指定文件。
	// _ = os.MkdirAll("excel", os.ModePerm) 目录授权
	filename := util.GetPath("book_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx")
	if err = f.SaveAs(filename); err != nil {
		fmt.Println(fmt.Sprintf("保存文件失败,错误:%s", err))
		return
	}
}
