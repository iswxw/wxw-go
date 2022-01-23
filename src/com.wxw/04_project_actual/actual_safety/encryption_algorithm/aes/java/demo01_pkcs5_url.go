/*
@Time : 2022/1/22 20:49
@Author : weixiaowei
@File : demo04_java
*/
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {

	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	content := "weixiaowei@qoogle.com"
	keyValue := "0123456789ABCDEF"
	key := []byte(keyValue)
	result, err := AesEncrypt([]byte(content), key)
	if err != nil {
		panic(err)
	}

	// 加密后： xBtQod-SPFDn0WVgbxa1lAwoUqffgf5nB_O4e9RO3PY
	fmt.Println("加密后：", base64.RawURLEncoding.EncodeToString(result))

	decodeString, _ := base64.RawURLEncoding.DecodeString("w0xFzL8zZVcpZq_KHzTmISwibcKpzwh6FseTq2KP8Xg")
	origData, err := AesDecrypt(decodeString, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(origData))
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding01(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding01(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//func ZeroPadding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{0}, padding)
//	return append(ciphertext, padtext...)
//}
//
//func ZeroUnPadding(origData []byte) []byte {
//	length := len(origData)
//	unpadding := int(origData[length-1])
//	return origData[:(length - unpadding)]
//}

//补码
func PKCS7Padding01(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS7UnPadding01(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
