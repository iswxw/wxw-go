/*
@Time: 2021/12/5 0:40
@Author: wxw
@File: unit
*/
package main

import "strings"

//NewSplit 切割字符串
//example:
//abc,b=>[ac]
func NewSplit(str, sep string) (des []string) {
	index := strings.Index(str, sep)
	for index > -1 {
		sectionBefor := str[:index]
		des = append(des, sectionBefor)
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	//最后1
	des = append(des, str)
	return
}
