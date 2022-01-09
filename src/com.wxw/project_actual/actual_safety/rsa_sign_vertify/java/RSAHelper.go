/*
@Time: 2022/1/10 0:19
@Author: wxw
@File: RSAHelper
#Link: https://blog.csdn.net/igoqhb/article/details/19832259
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
)

var decrypted string

func init() {
	flag.StringVar(&decrypted, "d", "", "加密过的数据")
	flag.Parse()
}

func main() {
	// 测试java生成的公钥和私钥加密和解密
	var data []byte
	var err error
	if decrypted != "" {
		data, err = base64.StdEncoding.DecodeString(decrypted)
		if err != nil {
			panic(err)
		}
	} else {
		//利用客户端传来的公钥加密有效信息
		data, err = RsaEncrypt([]byte("polaris@studygolang.com"), publicKey_by_java)
		if err != nil {
			panic(err)
		}
		fmt.Println("rsa encrypt(base64) : \n" + base64.StdEncoding.EncodeToString(data))
	}
	data, err = base64.StdEncoding.DecodeString("W/lT8KrQLB7Pj3sXI0sAwCyFjvCnVr/tUlgYHqUx3L8dZ5+sYMJqtatsas3Ks5qVmdMFBIKg4tZiA0WsqOQgt36z/xudRBvHxLIFSowpO4xjcym4vBaWnUiYEzJDed7jbSwaPHQTSLinqclSxbh32TKTJ9dFmihD2/vp0bfyt/k=")

	origData, err := RsaDecrypt(data, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("rsa decrypt(base64):\n" + string(origData))
}

// 公钥和私钥可以从文件中读取

// go 私钥
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

// java 生成的私钥
var privateKey_by_java = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIBNwIBADANBgkqhkiG9w0BAQEFAASCASEwggEdAgEAAoGBAOZvZteFx96KiyfGz
RsPXuKxqV6XJLb/TEpB0lRVTIx237/kieb9Lm6kMWuWIyzeL3OPc+0lY7vLNF+IGe
gM2U59Oo81joMj0VeeEOCNAwmGnVluGXt3vgTrcIFVNToN6hyp/L8oKI+NCelPOWf
z0ybt4TdZu+Wa9feY7u4Pb+B5AgEAAoGBAL5mmBxGzwIDib2hF0JfrfA0ChU9X7nR
MrE8t9S08l4xrul4pbV1x1LmWmtiD8h4Ac9DXe858LFv0uOIqpdBXp9ZMyoqBC97L
wDboutzt6OcXQ3hMVTOszn9cFFIf6JXaLz8HgocqAHTLVM4LwmyZNbGAyX/vja9BX
jVtUQdVxt9AgEAAgEAAgEAAgEAAgEA
-----END RSA PRIVATE KEY-----
`)

// go 公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

// java 公钥
var publicKey_by_java = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDmb2bXhcfeiosnxs0bD17isalel
yS2/0xKQdJUVUyMdt+/5Inm/S5upDFrliMs3i9zj3PtJWO7yzRfiBnoDNlOfTqPNY
6DI9FXnhDgjQMJhp1Zbhl7d74E63CBVTU6Deocqfy/KCiPjQnpTzln89Mm7eE3Wbv
lmvX3mO7uD2/geQIDAQAB
-----END PUBLIC KEY-----
`)

//利用客户端传来的公钥(java语言产生)加密
func RsaEncrypt(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密由客户端（java）利用服务器（Go语言实现，即本程序）公钥加密的信息
func RsaDecrypt(ciphertext []byte, privKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
