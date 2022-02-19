/*
@Time : 2022/1/24 19:22
@Author : weixiaowei
@File : demo04
*/
package main

import (
	"crypto/aes"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
)

// 加密后： w0xFzL8zZVcpZq_KHzTmISwibcKpzwh6FseTq2KP8Xg
// weixiaowei@qoogle.com
func main() {
	// content := "maple"
	content := "weixiaowei@qoogle.com" // 原文
	keyValue := "0123456789ABCDEF"     // 加密串、sign

	result := AesEncryptECB([]byte(content), []byte(keyValue))
	fmt.Println("加密后：", base64.RawURLEncoding.EncodeToString(result))

	decryptECB := AesDecryptECB(result, []byte(keyValue))
	fmt.Println(string(decryptECB))
}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	key, _ = AesSha1prng(key, 128)
	cipher, _ := aes.NewCipher(key)
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}

func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	key, _ = AesSha1prng(key, 128)
	cipher, _ := aes.NewCipher(key)
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted
}

// 模拟 java SHA1PRNG 处理
func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashes := Sha1(Sha1(keyBytes))
	maxLen := len(hashes)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New(fmt.Sprintf("invalid length!"))
	}
	return hashes[0:realLen], nil
}

func Sha1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
