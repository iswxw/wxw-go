/*
@Time : 2022/1/24 00:21
@Author : weixiaowei
@File : num2letterlib
*/
package dto

//数字转换为字母
func I2A(num int) string {
	if num < 0 {
		return ""
	}
	letterList := LetterList()
	return letterList[num]
}

func LetterList() []string {
	letterList := []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
		"AA",
		"AB",
		"AC",
		"AD",
		"AE",
		"AF",
		"AG",
		"AH",
		"AI",
		"AJ",
		"AK",
		"AL",
		"AM",
		"AN",
		"AO",
		"AP",
		"AQ",
		"AR",
		"AS",
		"AT",
		"AU",
		"AV",
		"AW",
		"AX",
		"AY",
		"AZ",
	}
	return letterList
}
