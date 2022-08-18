// @Time : 2022/8/18 11:16
// @Author : xiaoweiwei
// @File : integration

package integration

import (
	"context"
	"src/com.wxw/project_actual/src/com.wxw/03_thirdparty/w05_excel/common/dto"
)

type IExcelService interface {
	WriteToXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error)  // 数据写到Excel文件
	ReadFromXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error) // 从Excel文件读取数据
}

func (e excelService) WriteToXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error) {
	//
	//f := excelize.NewFile()
	//index := f.NewSheet(consts.SheetName)
	//
	////  设置表头
	//f.SetSheetRow(consts.SheetName, "A1", &param.Headers)
	//
	//// 设置列宽
	//for key, value := range param.ColumnWidths {
	//	f.SetColWidth(consts.SheetName, key, key, value)
	//}
	//
	//

	return nil, nil
}

func (e excelService) ReadFromXlsxFile(ctx context.Context, param *dto.ParamWriteExcel) ([]byte, error) {
	panic("implement me")
}

// 内部结构
type excelService struct {
}

var instance = &excelService{}

func NewExcelService() IExcelService {
	return instance
}
