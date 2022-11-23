/*
@Time: 2022/1/9 14:23
@Author: wxw
@File: EncryptUtil
*/
package utils

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

const TmpPath = "/docs/tmp/"

// [1] 生成RSA私钥和公钥，保存到文件中
// bits 证书大小
func GeneratorRsaPairKeys(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create(PemPath("private.pem"))
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(PemPath("public.pem"))
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)
}

// [2] 公钥加密
// plainText 要加密的数据
// path 公钥匙文件地址
func RSAEncrypt(plainText []byte, publicPath string) []byte {
	//打开文件
	currentKey, _ := ReadFileKey(publicPath)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(currentKey)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

// 读取秘钥内容
func ReadFileKey(filePath string) (key []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("open file err: ", err)
		return nil, err
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	return block.Bytes, nil
}

// 公钥加密-分段
func RSAEncryptBlock(plainText []byte, publicPath string) (bytesEncrypt []byte, err error) {
	//打开文件
	currentKey, _ := ReadFileKey(publicPath)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(currentKey)
	if err != nil {
		return
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	keySize, srcSize := publicKey.Size(), len(plainText)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt = buffer.Bytes()
	return
}

// [3] 私钥解密
// cipherText 需要解密的byte数据
// path 私钥文件路径
func RSADecrypt(cipherText []byte, privatePath string) []byte {
	//打开文件
	currentKey, _ := ReadFileKey(privatePath)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(currentKey)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}

// 私钥解密-分段
func RSADecryptBlock(cipherText []byte, privatePath string) (bytesDecrypt []byte, err error) {
	//打开文件
	currentKey, _ := ReadFileKey(privatePath)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(currentKey)
	if err != nil {
		return
	}
	keySize := privateKey.Size()
	srcSize := len(cipherText)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt = buffer.Bytes()
	return
}

// 私钥 签名
func Sign(signData []byte, privatePath string) []byte {
	//消息先进行Hash处理
	h := md5.New()
	h.Write(signData)
	hashed := h.Sum(nil)
	//签名
	//打开文件
	currentKey, _ := ReadFileKey(privatePath)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(currentKey)
	if err != nil {
		panic(err)
	}
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, e := rsa.SignPSS(rand.Reader, privateKey, crypto.MD5, hashed, opts)
	if e != nil {
		log.Println("sign err：", e)
	}
	return sig
}

// 公钥 验签
func Verify(signData []byte, publicPath string, sign []byte) bool {
	//打开文件
	currentKey, _ := ReadFileKey(publicPath)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(currentKey)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//消息先进行Hash处理
	h := md5.New()
	h.Write(signData)
	hashed := h.Sum(nil)

	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	if e := rsa.VerifyPSS(publicKey, crypto.MD5, hashed, sign, opts); e != nil {
		return false
	}
	return true
}

// 获取文件路径
func PemPath(fileName string) string {
	currentDir, _ := os.Getwd()
	return fmt.Sprintf("%s%s%s", currentDir, TmpPath, fileName)
}
