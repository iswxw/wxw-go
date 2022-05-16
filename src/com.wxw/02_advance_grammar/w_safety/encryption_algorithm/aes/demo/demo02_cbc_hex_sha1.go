/*
@Time : 2022/1/24 22:52
@Author : weixiaowei
@File : demo02_cbc_sha1
*/
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
)

// https://blog.csdn.net/wade3015/article/details/84454836
func main() {
	orig := "weixiaowei@qoogle.com" // 原文
	key := "0123456789ABCDEF"       // 加密串、sign
	fmt.Println("原文：", orig)

	encryptCode := AesEncrypt02(orig, key)
	fmt.Println("加密后：", encryptCode)
	decryptCode := AesDecrypt02(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}

// AES 加密
func AesEncrypt02(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)

	ck := []byte(key)
	k, _ := AesSha1prng(ck, 128)
	// 分组秘钥 // NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS5Padding02(origData, blockSize)
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
func AesDecrypt02(crypto string, key string) string {
	// 转成字节数组
	cryptoByte, _ := base64.RawURLEncoding.DecodeString(crypto)
	ck := []byte(key)
	k, _ := AesSha1prng(ck, 128)
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
	orig = PKCS5UnPadding02(orig)
	return string(orig)
}

//补码
func PKCS5Padding02(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS5UnPadding02(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashes := SHA1(SHA1(keyBytes))
	maxLen := len(hashes)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New(fmt.Sprintf("Not Support %d, Only Support Lower then %d [% x]", realLen, maxLen, hashes))
	}
	return hashes[0:realLen], nil
}

func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
