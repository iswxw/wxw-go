// @Time : 2023/3/20 17:57
// @Author : xiaoweiwei
// @File : consts

package java

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
)

//=============================枚举常量==================================

// Hash for crypto Hash
type Hash uint

const (
	MD5 Hash = iota
	SHA1
	SHA224
	SHA256
	SHA384
	SHA512
	SHA512_224
	SHA512_256
)

// Encode defines the type of bytes encoded to string
type Encode uint

const (
	String Encode = iota
	HEX
	Base64
)

// Secret defines the private key type
type Secret uint

const (
	PKCS1 Secret = iota
	PKCS8
)

//=============================结构体字段=================================

// RsaCrypt 加密结构体
type RsaCrypt struct {
	secretInfo RSASecret
}

// RSASecret 秘钥结构体
type RSASecret struct {
	PublicKey          string
	PublicKeyDataType  Encode
	PrivateKey         string
	PrivateKeyDataType Encode
	PrivateKeyType     Secret
}

//=============================通用方法==================================

// DecodeString decodes string data to bytes in designed encoded type
func DecodeString(data string, encodedType Encode) ([]byte, error) {
	var keyDecoded []byte
	var err error
	switch encodedType {
	case String:
		keyDecoded = []byte(data)
	case HEX:
		keyDecoded, err = hex.DecodeString(data)
	case Base64:
		keyDecoded, err = base64.StdEncoding.DecodeString(data)
	default:
		return keyDecoded, fmt.Errorf("secretInfo PublicKeyDataType unsupport")
	}
	return keyDecoded, err
}

// EncodeToString encodes data to string with encode type
func EncodeToString(data []byte, encodeType Encode) (string, error) {
	switch encodeType {
	case HEX:
		return hex.EncodeToString(data), nil
	case Base64:
		return base64.StdEncoding.EncodeToString(data), nil
	case String:
		return string(data), nil
	default:
		return "", fmt.Errorf("secretInfo OutputType unsupport")
	}
}

// split 拆分
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

// ParsePrivateKey parses private key bytes to rsa privateKey
func ParsePrivateKey(privateKeyDecoded []byte, keyType Secret) (*rsa.PrivateKey, error) {
	switch keyType {
	case PKCS1:
		return x509.ParsePKCS1PrivateKey(privateKeyDecoded)
	case PKCS8:
		keyParsed, err := x509.ParsePKCS8PrivateKey(privateKeyDecoded)
		return keyParsed.(*rsa.PrivateKey), err
	default:
		return &rsa.PrivateKey{}, fmt.Errorf("secretInfo PrivateKeyDataType unsupport")
	}
}

// GetHash gets the crypto hash type & hashed data in different hash type
func GetHash(data []byte, hashType Hash) (h crypto.Hash, hashed []byte, err error) {
	nh, h := GetHashFunc(hashType)
	hh := nh()
	if _, err = hh.Write(data); err != nil {
		return
	}
	hashed = hh.Sum(nil)
	return
}

// GetHashFunc gets the crypto hash func & type in different hash type
func GetHashFunc(hashType Hash) (f func() hash.Hash, h crypto.Hash) {
	switch hashType {
	case SHA1:
		f = sha1.New
		h = crypto.SHA1
	case SHA224:
		f = sha256.New224
		h = crypto.SHA224
	case SHA256:
		f = sha256.New
		h = crypto.SHA256
	case SHA384:
		f = sha512.New384
		h = crypto.SHA384
	case SHA512:
		f = sha512.New
		h = crypto.SHA512
	case SHA512_224:
		f = sha512.New512_224
		h = crypto.SHA512_224
	case SHA512_256:
		f = sha512.New512_256
		h = crypto.SHA512_256
	case MD5:
		f = md5.New
		h = crypto.MD5
	default:
		panic(any("not support hashType"))
	}
	return
}
