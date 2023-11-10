/*
@Time : 2022/1/20 23:32
@Author : weixiaowei
@File : main
*/
package main

import "encoding/base64"

const (
	base64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"
)

var coder = base64.NewEncoding(base64Table)

// https://blog.csdn.net/wade3015/article/details/84454836
func Base64Encode(src []byte) []byte { //编码
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) { //解码
	return coder.DecodeString(string(src))
}
