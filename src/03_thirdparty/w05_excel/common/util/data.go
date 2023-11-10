/*
@Time : 2022/1/26 00:29
@Author : weixiaowei
@File : data
*/
package util

import (
	dto2 "src/com.wxw/project_actual/src/03_thirdparty/w05_excel/common/dto"
)

func GetUsers() []dto2.User {
	return []dto2.User{
		{
			Id:        1,
			Name:      "北京",
			Amount:    123,
			Price:     "999",
			UnitPrice: 123,
		},
		{
			Id:        2,
			Name:      "上海",
			Amount:    456,
			Price:     "9999",
			UnitPrice: 445,
		},
	}
}
