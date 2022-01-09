/*
@Time: 2022/1/9 16:24
@Author: wxw
@File: demo_rsa_partition
@Link: https://blog.csdn.net/u013650708/article/details/85337520
*/
package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"log"
)

// 分段加密解密
func main() {
	text := "我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅我不管我就是最帅"
	log.Println("加密前：\t", text)

	md5 := MD5([]byte(text))
	//MD5打印为16进制字符串
	log.Println("MD5打印为16进制字符串：", hex.EncodeToString(md5))

	//RSA的内容使用base64打印
	privateKey, publicKey, _ := GenRSAKey(1024)
	log.Println("rsa私钥:\t", base64.StdEncoding.EncodeToString(privateKey))
	log.Println("rsa公钥:\t", base64.StdEncoding.EncodeToString(publicKey))

	plains, err := RsaEncryptBlock([]byte(text), publicKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("加密后：\t", base64.StdEncoding.EncodeToString(plains))

	cipherText, err := RsaDecryptBlock(plains, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("解密后：\t", string(cipherText))
}

// MD5 hash
func MD5(src []byte) []byte {
	hash := md5.New()
	hash.Write(src)
	return hash.Sum(nil)
}

/**
 * 生成RSA密钥对
 */
func GenRSAKey(size int) (privateKeyBytes, publicKeyBytes []byte, err error) {
	//生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return
	}
	privateKeyBytes = x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes = x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	return
}

/**
 * 公钥加密
 */
func RsaEncrypt(src, publicKeyByte []byte) (bytes []byte, err error) {
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyByte)
	if err != nil {
		return
	}
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
}

/**
 * 公钥加密-分段
 */
func RsaEncryptBlock(src, publicKeyByte []byte) (bytesEncrypt []byte, err error) {
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyByte)
	if err != nil {
		return
	}
	keySize, srcSize := publicKey.Size(), len(src)
	log.Println("密钥长度：", keySize, "\t明文长度：\t", srcSize)
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
 * 私钥解密
 */
func RsaDecrypt(src, privateKeyBytes []byte) (bytesDecrypt []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
}

/**
 * 私钥解密-分段
 */
func RsaDecryptBlock(src, privateKeyBytes []byte) (bytesDecrypt []byte, err error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return
	}
	keySize := privateKey.Size()
	srcSize := len(src)
	log.Println("密钥长度：", keySize, "\t密文长度：\t", srcSize)
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
