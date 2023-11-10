// @Time : 2023/3/20 18:00
// @Author : xiaoweiwei
// @File : rsa

package java

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

func NewRSACrypt(secretInfo RSASecret) *RsaCrypt {
	return &RsaCrypt{secretInfo: secretInfo}
}

// Encrypt 加密
func (rc *RsaCrypt) Encrypt(src string, outputDataType Encode) (dst string, err error) {
	secretInfo := rc.secretInfo
	if secretInfo.PublicKey == "" {
		return "", fmt.Errorf("secretInfo PublicKey can't be empty")
	}
	pubKeyDecoded, err := DecodeString(secretInfo.PublicKey, secretInfo.PublicKeyDataType)
	if err != nil {
		return
	}
	pubKey, err := x509.ParsePKIXPublicKey(pubKeyDecoded)
	if err != nil {
		return
	}
	publicKey := pubKey.(*rsa.PublicKey)
	partLen := publicKey.N.BitLen()/8 - 11
	chunks := split([]byte(src), partLen)

	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		byteData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(byteData)
	}

	return EncodeToString(buffer.Bytes(), outputDataType)
}

// Decrypt 解密
func (rc *RsaCrypt) Decrypt(src string, srcType Encode) (dst string, err error) {
	secretInfo := rc.secretInfo
	if secretInfo.PrivateKey == "" {
		return "", fmt.Errorf("secretInfo PrivateKey can't be empty")
	}
	privateKeyDecoded, err := DecodeString(secretInfo.PrivateKey, secretInfo.PrivateKeyDataType)
	if err != nil {
		return
	}
	prvKey, err := ParsePrivateKey(privateKeyDecoded, secretInfo.PrivateKeyType)
	if err != nil {
		return
	}
	partLen := prvKey.N.BitLen() / 8
	decodeData, err := DecodeString(src, srcType)
	chunks := split(decodeData, partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(decrypted)
	}
	return buffer.String(), nil

}

// Sign 生成签名
func (rc *RsaCrypt) Sign(src string, hashType Hash, outputDataType Encode) (dst string, err error) {
	secretInfo := rc.secretInfo
	if secretInfo.PrivateKey == "" {
		return "", fmt.Errorf("secretInfo PrivateKey can't be empty")
	}
	privateKeyDecoded, err := DecodeString(secretInfo.PrivateKey, secretInfo.PrivateKeyDataType)
	if err != nil {
		return
	}
	prvKey, err := ParsePrivateKey(privateKeyDecoded, secretInfo.PrivateKeyType)
	if err != nil {
		return
	}
	cryptoHash, hashed, err := GetHash([]byte(src), hashType)
	if err != nil {
		return
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, prvKey, cryptoHash, hashed)
	if err != nil {
		return
	}
	return EncodeToString(signature, outputDataType)
}

// VerifySign 校验签名
func (rc *RsaCrypt) VerifySign(src string, hashType Hash, signedData string, signDataType Encode) (bool, error) {
	secretInfo := rc.secretInfo
	if secretInfo.PublicKey == "" {
		return false, fmt.Errorf("secretInfo PublicKey can't be empty")
	}
	publicKeyDecoded, err := DecodeString(secretInfo.PublicKey, secretInfo.PublicKeyDataType)
	if err != nil {
		return false, err
	}
	pubKey, err := x509.ParsePKIXPublicKey(publicKeyDecoded)
	if err != nil {
		return false, err
	}
	cryptoHash, hashed, err := GetHash([]byte(src), hashType)
	if err != nil {
		return false, err
	}
	signDecoded, err := DecodeString(signedData, signDataType)
	if err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), cryptoHash, hashed, signDecoded); err != nil {
		return false, err
	}
	return true, nil
}
