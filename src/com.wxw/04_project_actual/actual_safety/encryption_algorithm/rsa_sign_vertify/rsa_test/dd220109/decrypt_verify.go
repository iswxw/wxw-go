/*
@Time : 2022/1/9 21:12
@Author : wxw
@File : decrypt_verify 合作方：私钥 解密 + 签名
*/
package dd220109

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
)

// 主办方：测试 签名 和 解密

/**
 * 私钥解密-分段
 */
func DecryptBlock(src, privateKeyBytes []byte) (bytesDecrypt []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return
	}
	keySize := privateKey.Size()
	srcSize := len(src)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt = buffer.Bytes()
	return
}

/**
 * 公钥 验签
 */
func Verify(src, publicKeyByte []byte, sign []byte) (isSuccess bool, err error) {

	// 公钥
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyByte)
	if err != nil {
		return false, err
	}

	// MD5
	hashed := MD5(src)

	// 验签参数
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}

	// 验签
	if err := rsa.VerifyPSS(publicKey, crypto.MD5, hashed, sign, opts); err != nil {
		return false, err
	}
	return true, nil
}
