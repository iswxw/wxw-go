/*
@Time : 2022/1/22 02:44
@Author : weixiaowei
@File : demo03_base64_url
*/
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 对接三方可以使用，有加密解密拆分交互场景
// 1. 处理了由于base64编码在url请求传输导致 解码失败的问题(https://blog.csdn.net/u014270740/article/details/91038606)
func main() {

	// 加密
	aesKey := "0123456789ABCDEF"
	content := "weixiaowei@qoogle.com"
	encrypt03 := AesEncrypt03(content, aesKey)
	fmt.Println("加密后：", encrypt03)

	decrypt03 := AesDecrypt03(encrypt03, aesKey)
	fmt.Println("解密后：", decrypt03)

}

// AES 加密
func AesEncrypt03(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding03(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	// return base64.StdEncoding.EncodeToString(cryted)
	return base64.RawURLEncoding.EncodeToString(cryted)
}

// AES 解密
func AesDecrypt03(crypto string, key string) string {
	// 转成字节数组
	cryptoByte, _ := base64.RawURLEncoding.DecodeString(crypto)
	// cryptoByte, _ := base64.StdEncoding.DecodeString(crypto) // url请求参数中会发生转码，导致解析失败
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryptoByte))
	// 解密
	blockMode.CryptBlocks(orig, cryptoByte)
	// 去补全码
	orig = PKCS7UnPadding03(orig)
	return string(orig)
}

//补码
func PKCS7Padding03(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS7UnPadding03(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
