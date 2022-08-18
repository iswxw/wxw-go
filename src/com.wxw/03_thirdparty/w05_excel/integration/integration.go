// @Time : 2022/8/18 11:16
// @Author : xiaoweiwei
// @File : integration

package integration

import (
	"context"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"src/com.wxw/project_actual/src/com.wxw/03_thirdparty/w05_excel/common/consts"
	"src/com.wxw/project_actual/src/com.wxw/03_thirdparty/w05_excel/common/dto"
)

type IExcelService interface {
	WriteToXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error)  // 数据写到Excel文件
	ReadFromXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error) // 从Excel文件读取数据
}

func (e excelService) WriteToXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error) {

	f := excelize.NewFile()

	index := f.NewSheet(consts.SheetName)

	//  设置表头
	f.SetSheetRow(consts.SheetName, "A1", &param.Headers)

	// 设置列宽
	for key, value := range param.ColumnWidths {
		f.SetColWidth(consts.SheetName, key, key, value)
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// TODO 记录日志
	log.Println("param = ", param)

	// 数据写入文件缓存
	writeToBuffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return writeToBuffer.Bytes(), nil
}

func (e excelService) ReadFromXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error) {

	return nil, nil
}

// 内部结构
type excelService struct {
}

var instance = &excelService{}

func NewExcelService() IExcelService {
	return instance
}
