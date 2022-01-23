/*
@Time : 2022/1/23 00:44
@Author : weixiaowei
@File : demo05_generator_key
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

	key := []byte("0123456789ABCDEF")
	result, err := AesEncrypt05([]byte(content), key)
	if err != nil {
		panic(err)
	}

	// 加密后： xBtQod-SPFDn0WVgbxa1lAwoUqffgf5nB_O4e9RO3PY
	fmt.Println("加密后：", base64.RawURLEncoding.EncodeToString(result))
	decodeString, _ := base64.RawURLEncoding.DecodeString("xBtQod-SPFDn0WVgbxa1lAwoUqffgf5nB_O4e9RO3PY")
	origData, err := AesDecrypt05(decodeString, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密后：", string(origData))
}

func AesEncrypt05(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding05(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt05(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding05(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//补码
func PKCS7Padding05(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//去码
func PKCS7UnPadding05(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

/**
 *
 * AES CBC 256 位加密
 * @param content 加密内容字节数组
 * @param keyBytes 加密字节数组
 * @param iv 加密向量字节数组
 * @param encryptLength 仅支持 128、256 长度
 * @return 解密后字节内容
 */
/*public static byte[] encryptAESCBC( byte[] content, byte[] keyBytes,byte[] iv, int encryptLength ){

KeyGenerator	keyGenerator	= KeyGenerator.getInstance( "AES" );
SecureRandom	secureRandom	= SecureRandom.getInstance( "SHA1PRNG" );
secureRandom.setSeed( keyBytes );
keyGenerator.init( encryptLength, secureRandom );
SecretKey	key	= keyGenerator.generateKey();
// 以上应该是使用任意长度的 keyBytes，生成指定长度为 encryptLength 的 key
// 后面再使用 AES/CBC/PKCS5Padding 算法，对 content 加密，加密角度是上面生成的固定长度的 key
// 我的主要疑问就在这里，  如何用 golang 生成与 keyGenerator.generateKey() 一样的 key 呢？

Cipher		cipher	= Cipher.getInstance( "AES/CBC/PKCS5Padding" );
cipher.init( Cipher.ENCRYPT_MODE, key, new IvParameterSpec( iv ) );
byte[] result = cipher.doFinal( content );
return(result);
}*/
