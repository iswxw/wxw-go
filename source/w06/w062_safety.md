## 工程实践之安全

### 基础回顾

基础知识可以查看 [Java 签名和验签相关文章](https://github.com/iswxw/wxw-distributed/tree/dev-wxw/cloud-safety) ，在尝试RSA加/解密的时候，发现go标准库中仅有"公钥加密，私钥解密"，而没有“私钥加密、公钥解密”。经过考虑，我认为GO的开发者是故意这样设计的，原因如下：

1. 非对称加密相比对称加密的好处就是：私钥自己保留，公钥公布出去，公钥加密后只有私钥能解开，私钥加密后只有公钥能解开。
2. 如果仅有一对密钥，与对称加密区别就不大了。假如你是服务提供方，使用私钥进行加密后，接入方使用你提供的公钥进行解密，一旦这个公钥泄漏，带来的后果和对称加密密钥泄漏是一样的。只有双方互换公钥（均使用对方公钥加密，己方私钥解密），才能充分发挥非对称加密的优势。

当然，有第三方库支持“私钥加密、公钥解密”的，有兴趣的伙伴可自行百度。

关于go加密和解密实现如下：

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

- 通过 **签名和验签** 避免 信息被篡改

- 通过 **分段加密和解密** 解决 待解密解密信息太长导致解密和解密失败的问题

  **RSA加密时，明文长度>（密钥长度-padding长度）时需要进行分段。PKCS1填充长度为11**  

```go
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
```

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