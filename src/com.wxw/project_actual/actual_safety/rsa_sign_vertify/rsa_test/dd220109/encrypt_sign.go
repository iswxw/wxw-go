/*
@Time : 2022/1/9 21:11
@Author : wxw
@File : encrypt_sign 业务方：公钥 加密+验签
*/
package dd220109

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
)

// 业务方：测试加密 和 签名

/**
 * 公钥加密-分段
 */
func EncryptBlock(src, publicKeyByte []byte) (bytesEncrypt []byte, err error) {
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	if err != nil {
		return nil, err
	}
	keySize, srcSize := publicKey.Size(), len(src)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt = buffer.Bytes()
	return
}

/**
 * 私钥 签名
 */
func Sign(src, privateKeyBytes []byte) (sign []byte, err error) {

	// 获取私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	// 获取hashed
	hashed := MD5(src)

	// 签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sign, err = rsa.SignPSS(rand.Reader, privateKey, crypto.MD5, hashed, opts)
	if err != nil {
		log.Println("签名错误：", err)
		return nil, err
	}
	return sign, nil
}

// MD5 hash
func MD5(src []byte) []byte {
	hash := md5.New()
	hash.Write(src)
	return hash.Sum(nil)
}
