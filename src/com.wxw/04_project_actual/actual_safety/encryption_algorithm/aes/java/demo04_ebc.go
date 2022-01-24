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

func main() {
	content := "maple"
	keyValue := "0123456789ABCDEF"
	// key,_ := AesSha1prng([]byte(keyValue),128)

	result := AesEncryptECB([]byte(content), []byte(keyValue))
	// 加密后： xBtQod-SPFDn0WVgbxa1lAwoUqffgf5nB_O4e9RO3PY
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

// content：test123
// encryptKey：123456
// 加密结果为：668C826342B8703D86E8BBF404610499
// 此时就和 java 结果相对应了，解密也一样对 key 加一步处理就行
//func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
//	key, _ = AesSha1prng(key, 128) // 比示例一多出这一步
//	//cipher, _ := aes.NewCipher(generateKey(key))
//	cipher, _ := aes.NewCipher(key)
//	length := (len(origData) + aes.BlockSize) / aes.BlockSize
//	plain := make([]byte, length*aes.BlockSize)
//	copy(plain, origData)
//	pad := byte(len(plain) - len(origData))
//	for i := len(origData); i < len(plain); i++ {
//		plain[i] = pad
//	}
//	encrypted = make([]byte, len(plain))
//	// 分组分块加密
//	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
//		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
//	}
//
//	return encrypted
//}

// 模拟 java SHA1PRNG 处理
func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashs := Sha1(Sha1(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New("invalid length!")
	}
	return hashs[0:realLen], nil
}

func Sha1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
