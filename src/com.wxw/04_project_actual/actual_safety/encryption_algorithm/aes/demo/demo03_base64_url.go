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
)

// 对接三方可以使用，有加密解密拆分交互场景
// 1. 处理了由于base64编码在url请求传输导致 解码失败的问题(https://blog.csdn.net/u014270740/article/details/91038606)
func main() {
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
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.RawURLEncoding.EncodeToString(cryted)
}

// AES 解密
func AesDecrypt03(crypto string, key string) string {
	// 转成字节数组
	cryptoByte, _ := base64.RawURLEncoding.DecodeString(crypto)
	// cryptoByte, _ := base64.StdEncoding.DecodeString(crypto) //
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
	orig = PKCS7UnPadding(orig)
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
