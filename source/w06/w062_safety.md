## 工程实践之安全

### 基础回顾

基础知识可以查看 [Java 签名和验签相关文章](https://github.com/iswxw/wxw-distributed/tree/dev-wxw/cloud-safety) ，关于go加密和解密实现如下：

#### 1. 实现加密解密

```go
package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// rsa加密解密 签名验签
func main() {
	//生成私钥
	pri, e := rsa.GenerateKey(rand.Reader, 1024)
	if e != nil {
		fmt.Println(e)
	}

	//根据私钥产生公钥
	pub := &pri.PublicKey

	//明文
	plaintext := []byte("Hello world")

	//加密生成密文
	fmt.Printf("加密前： %q\n", plaintext)
	ciphertext, e := rsa.EncryptOAEP(md5.New(), rand.Reader, pub, plaintext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("加密后：%x\n", ciphertext)

	//解密得到明文
	plaintext, e = rsa.DecryptOAEP(md5.New(), rand.Reader, pri, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("解密后： %q\n", plaintext)

	//消息先进行Hash处理
	h := md5.New()
	h.Write(plaintext)
	hashed := h.Sum(nil)
	fmt.Printf("%q MD5 Hashed:\t%x\n", plaintext, hashed)

	//签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, e := rsa.SignPSS(rand.Reader, pri, crypto.MD5, hashed, opts)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("签名: \t%x\n", sig)

	//认证
	fmt.Printf("验证结果:")
	if e := rsa.VerifyPSS(pub, crypto.MD5, hashed, sig, opts); e != nil {
		fmt.Println("失败:", e)
	} else {
		fmt.Println("成功.")
	}
}
```

#### 2. 问题和方案

- 通过签名和验签 避免 信息被篡改
- 通过分段加密和解密解决 带解析信息太长问题



### 辅助工具

#### 1. RAS 证书生成

- openssl方式生成

  ```bash
  # 生成私钥
  openssl genrsa -out rsa_private_key.pem 1024
  
  # 生成公钥
  openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
  ```

- Go 代码生成

  ```go
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
  
  // 获取文件路径
  func PemPath(fileName string) string {
  	currentDir, _ := os.Getwd()
  	return fmt.Sprintf("%s%s%s",currentDir, consts.TmpPath, fileName)
  }
  
  // 文件路径
  const (
  	TmpPath = "/docs/tmp/"
  )
  ```

  

---

相关资料

1. [Java 公钥加密私钥解密实现理论](https://github.com/iswxw/wxw-distributed/tree/dev-wxw/cloud-safety) 
2. [Golang 实现RSA加密解密](https://www.cnblogs.com/zhichaoma/p/12516715.html)  
3. [Golang RSA实现加密解密和签名与验签](https://www.jianshu.com/p/0d4954aad89f) 