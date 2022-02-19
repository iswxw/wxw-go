/*
@Time : 2022/1/24 23:05
@Author : weixiaowei
@File : demo01_java_cbc_base64
*/
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

// Wqm8-LnVkjfurtZT5ntvJ23fwA4H7dd6frJegphx2v0
func main() {
	content := "weixiaowei@qoogle.com" // 原文
	key := "0123456789ABCDEF"          // 加密串、sign

	result, err := AesEncrypt01(content, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密后：", result)

	origData, err := AesDecrypt01(result, key)
	if err != nil {
		log.Println("err:", err)
	}
	fmt.Println("解密后：", origData)
}

// AES 加密
func AesEncrypt01(orig string, key string) (string, error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS5Padding01(origData, blockSize)
	// 创建数组
	crypted := make([]byte, len(origData))
	err := cryptBlocks01(block, origData, crypted)
	if err != nil {
		return "", err
	}
	//return base64.StdEncoding.EncodeToString(crypted), nil
	return base64.RawURLEncoding.EncodeToString(crypted), nil
}

// AES 解密
func AesDecrypt01(crypto string, key string) (string, error) {
	// 转成字节数组
	crypted, _ := base64.RawURLEncoding.DecodeString(crypto)
	// crypted, _ := base64.StdEncoding.DecodeString(crypto)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 创建数组
	decrypted := make([]byte, len(crypted))
	// 解密处理
	if err := decryptBlocks01(block, crypted, decrypted); err != nil {
		return "", err
	}
	// 去补全码
	decrypted = PKCS5UnPadding01(decrypted)
	return string(decrypted), nil
}

//补码
func PKCS5Padding01(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS5UnPadding01(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func cryptBlocks01(cb cipher.Block, src, dst []byte) error {

	if len(src)%cb.BlockSize() != 0 {
		return fmt.Errorf("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		return fmt.Errorf("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		cb.Encrypt(dst, src[:cb.BlockSize()])
		src = src[cb.BlockSize():]
		dst = dst[cb.BlockSize():]
	}
	return nil
}

func decryptBlocks01(cb cipher.Block, src, dst []byte) error {
	if len(src)%cb.BlockSize() != 0 {
		return fmt.Errorf("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		return fmt.Errorf("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		cb.Decrypt(dst, src[:cb.BlockSize()])
		src = src[cb.BlockSize():]
		dst = dst[cb.BlockSize():]
	}
	return nil
}
