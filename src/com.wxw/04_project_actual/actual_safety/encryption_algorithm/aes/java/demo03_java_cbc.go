/*
@Time : 2022/1/24 13:53
@Author : weixiaowei
@File : demo02_java
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

func main() {
	content := "maple"
	keyValue := "0123456789ABCDEF"
	key, _ := AESSHA1PRNG03([]byte(keyValue), 128)

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

// AES 加密
func AesEncrypt03(orig string, key string) (string, error) {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS5Padding03(origData, blockSize)
	// 创建数组
	crypted := make([]byte, len(origData))
	err := cryptBlocks03(block, origData, crypted)
	if err != nil {
		return "", err
	}
	//// 加密模式
	//blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	//// 加密
	//blockMode.CryptBlocks(cryted, origData)
	// return base64.RawURLEncoding.EncodeToString(crypted), nil
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AES 解密
func AesDecrypt03(crypto string, key string) (string, error) {
	// 转成字节数组
	crypted, _ := base64.RawURLEncoding.DecodeString(crypto)
	// crypted, _ := base64.StdEncoding.DecodeString(crypto)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 创建数组
	decrypted := make([]byte, len(crypted))
	// 解密处理
	if err := decryptBlocks03(block, crypted, decrypted); err != nil {
		return "", err
	}
	//// 加密模式
	//blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	//// 解密
	//blockMode.CryptBlocks(orig, cryptoByte)
	// 去补全码
	decrypted = PKCS5UnPadding03(decrypted)
	return string(decrypted), nil
}

//补码
func PKCS5Padding03(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS5UnPadding03(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func cryptBlocks03(cb cipher.Block, src, dst []byte) error {
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

func decryptBlocks03(cb cipher.Block, src, dst []byte) error {
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

func AESSHA1PRNG03(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashs := SHA1(SHA1(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New(fmt.Sprintf("Not Support %d, Only Support Lower then %d [% x]", realLen, maxLen, hashs))
	}

	return hashs[0:realLen], nil
}

func SHA103(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
