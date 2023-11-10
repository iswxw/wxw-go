// @Time : 2023/3/24 16:07
// @Author : xiaoweiwei
// @File : const

package cvalidator

// 正则表达式
const (
	// UpperNumber 正则：大写字母/数字
	// 保单号、发动机号都用这个正则
	UpperNumber = "^[A-Z0-9]+$"

	// FrameNo 车架号正则：长度为17位，大写字母和数字组成
	FrameNo = "^[A-Z0-9]{17}$"

	// ChineseCharacter 汉字正则
	ChineseCharacter = "^[\u4e00-\u9fa5]+$"

	// YearMonthDayRegexp YYYY-MM-DD时间格式
	YearMonthDayRegexp = "^((19|20)[0-9]{2})-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$"

	// YearMonthRegexp YYYY-MM时间格式
	YearMonthRegexp = "^((19|20)[0-9]{2})-(0[1-9]|1[012])$"
)
