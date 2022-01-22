## 工程实践之安全

### 基础回顾

在项目开发过程中，当操作一些用户的隐私信息，诸如密码、帐户密钥等数据时，往往需要加密后可以在网上传输。这时，需要一些高效地、简单易用的加密算法加密数据，然后把加密后的数据存入数据库或进行其他操作；当需要读取数据时，把加密后的数据取出来，再通过算法解密。

当前我们项目中常用的加解密的方式无非三种.

- 对称加密, 加解密都使用的是同一个密钥, 其中的代表就是AES、DES
- 非对加解密, 加解密使用不同的密钥, 其中的代表就是RSA
- 签名算法, 如MD5、SHA1、HMAC等, 主要用于验证，防止信息被修改, 如：文件校验、数字签名、鉴权协议

**Base64不是加密算法**，它是一种数据编码方式，虽然是可逆的，但是它的编码方式是公开的，无所谓加密。本文也对Base64编码方式做了简要介绍。

**优先级** 

```bash
AES>DES
```

### 加密算法

#### 1. AES

AES，即高级加密标准（Advanced Encryption Standard），是一个对称分组密码算法，旨在取代DES成为广泛使用的标准。AES中常见的有三种解决方案，分别为AES-128、AES-192和AES-256。

<img src="https://imgconvert.csdnimg.cn/aHR0cDovL2ltZy5ibG9nLmNzZG4ubmV0LzIwMTcwMjE5MDgyOTA5Njg4?x-oss-process=image/format,png" alt="加密流程图" style="zoom: 50%;" /> 



##### 1.1 AES基本结构

AES为分组密码，分组密码也就是把明文分成一组一组的，每组长度相等，每次加密一组数据，直到加密完整个明文。在AES标准规范中，分组长度只能是128位，也就是说，每个分组为16个字节（每个字节8位）。密钥的长度可以使用128位、192位或256位。密钥的长度不同，推荐加密轮数也不同，如下表所示：

| AES     | 密钥长度（32位比特字) | 分组长度(32位比特字) | 加密轮数 |
| ------- | --------------------- | -------------------- | -------- |
| AES-128 | 4                     | 4                    | 10       |
| AES-192 | 6                     | 4                    | 12       |
| AES-256 | 8                     | 4                    | 14       |

轮数在下面介绍，这里实现的是AES-128，也就是密钥的长度为128位，加密轮数为10轮。
上面说到，AES的加密公式为C = E(K,P)，在加密函数E中，会执行一个轮函数，并且执行10次这个轮函数，这个轮函数的前9次执行的操作是一样的，只有第10次有所不同。也就是说，一个明文分组会被加密10轮。AES的核心就是实现一轮中的所有操作。

```bash
128bit = 16字节
192bit = 24字节
256bit = 32字节

在go提供的官方接口中秘钥长度只支持16字节
```

##### 1.2 **AES加密和解密** 

AES加密过程涉及到4种操作：字节替代（SubBytes）、行移位（ShiftRows）、列混淆（MixColumns）和轮密钥加（AddRoundKey）。解密过程分别为对应的逆操作。由于每一步操作都是可逆的，按照相反的顺序进行解密即可恢复明文。加解密中每轮的密钥分别由初始密钥扩展得到。算法中16字节的明文、密文和轮密钥都以一个4x4的矩阵表示。
AES 有五种加密模式：

- 电码本模式（Electronic Codebook Book (ECB)）
- 密码分组链接模式（Cipher Block Chaining (CBC)）
- 计算器模式（Counter (CTR)）
- 密码反馈模式（Cipher FeedBack (CFB)）
- 输出反馈模式（Output FeedBack (OFB)）

 **加密模式价值分析**  

对称/[分组密码](https://baike.baidu.com/item/分组密码)一般分为流加密(如OFB、CFB等)和块加密(如ECB、CBC等)。对于流加密，需要将分组密码转化为流模式工作。对于块加密(或称分组加密)，如果要加密超过块大小的数据，就需要涉及填充和链加密模式。

- CBC(Cipher Block Chaining，加密块链)模式

  ```bash
  优点
    1.不容易主动攻击,安全性好于ECB,适合传输长度长的报文,是SSL、IPSec的标准。
  缺点：　
    1.不利于并行计算；　
    2.误差传递；　
    3.需要初始化向量IV
  ```

- ECB(Electronic Code Book电子密码本)模式

  ECB模式是最早采用和最简单的模式，它将加密的数据分成若干组，每组的大小跟加密[密钥](https://baike.baidu.com/item/密钥)长度相同，然后每组都用相同的密钥进行加密。

  ```bash
  优点:
    1.简单；　
    2.有利于并行计算；　
    3.误差不会被传送；　
  缺点:　
    1.不能隐藏明文的模式；　
    2.可能对明文进行主动攻击；
    
  因此，此模式适于加密小消息。
  ```

##### 1.3 AES 算法原理

**AES原理**：AES是对数据按128位，也就是16个字节进行分组进行加密的，每次对一组数据加密需要运行多轮。而输入密钥的长度可以为128、192和256位，也就是16个字节、24个字节和32个字节，如果用户输入的密钥长度不是这几种长度，也会补成这几种长度。无论输入密钥是多少字节，加密还是以16字节的数据一组来进行的，密钥长度的不同仅仅影响加密运行的轮数。

##### 1.4 答疑解惑

1. 三种填充模式的区别(PKCS7Padding/PKCS5Padding/ZeroPadding)

   ```bash
   某些加密算法要求明文需要按一定长度对齐，叫做块大小(BlockSize)，比如16字节，那么对于一段任意的数据，加密前需要对最后一个块填充到16 字节，解密后需要删除掉填充的数据。
       - ZeroPadding，数据长度不对齐时使用0填充，否则不填充。
       - PKCS7Padding，假设数据长度需要填充n(n>0)个字节才对齐，那么填充n个字节，每个字节都是n;如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小。
       - PKCS5Padding，PKCS7Padding的子集，块大小固定为8字节。
   由于使用PKCS7Padding/PKCS5Padding填充时，最后一个字节肯定为填充数据的长度，所以在解密后可以准确删除填充的数据，而使用ZeroPadding填充时，没办法区分真实数据与填充数据，所以只适合以\0结尾的字符串加解密。
   ```

   

相关文档

1. http://docscn.studygolang.com/pkg/crypto/aes/
2. [go aes 加密算法总结](https://blog.csdn.net/wade3015/article/details/84454836) 
3. [AES 百度百科](https://baike.baidu.com/item/%E9%AB%98%E7%BA%A7%E5%8A%A0%E5%AF%86%E6%A0%87%E5%87%86/468774?fromtitle=aes&fromid=5903&fr=aladdin) 
3. [AES 详解](https://github.com/matt-wu/AES) 

#### 2. DES 

DES是一种对称加密算法，又称为美国数据加密标准。DES加密时以64位分组对数据进行加密，加密和解密都使用的是同一个长度为64位的密钥，实际上只用到了其中的56位，密钥中的第8、16…64位用来作奇偶校验。DES有ECB（电子密码本）和CBC（加密块）等加密模式。
DES算法的安全性很高，目前除了穷举搜索破解外， 尚无更好的的办法来破解。其密钥长度越长，破解难度就越大。



相关文档

1. [go DES 加密算法总结](https://blog.csdn.net/wade3015/article/details/84454836)  

#### 3.RSA

基础知识可以查看 [Java 签名和验签相关文章](https://github.com/iswxw/wxw-distributed/tree/dev-wxw/cloud-safety) ，在尝试RSA加/解密的时候，发现go标准库中仅有"公钥加密，私钥解密"，而没有“私钥加密、公钥解密”。经过考虑，我认为GO的开发者是故意这样设计的，原因如下：

1. 非对称加密相比对称加密的好处就是：私钥自己保留，公钥公布出去，公钥加密后只有私钥能解开，私钥加密后只有公钥能解开。
2. 如果仅有一对密钥，与对称加密区别就不大了。假如你是服务提供方，使用私钥进行加密后，接入方使用你提供的公钥进行解密，一旦这个公钥泄漏，带来的后果和对称加密密钥泄漏是一样的。只有双方互换公钥（均使用对方公钥加密，己方私钥解密），才能充分发挥非对称加密的优势。

当然，有第三方库支持“私钥加密、公钥解密”的，有兴趣的伙伴可自行百度。

关于go加密和解密实现如下：

##### 3.1 实现加密解密

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

##### 3.2 问题和方案

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

#### 4. MD5

MD5的全称是Message-DigestAlgorithm 5，它可以把一个任意长度的字节数组转换成一个定长的整数，并且这种转换是不可逆的。对于任意长度的数据，转换后的MD5值长度是固定的，而且MD5的转换操作很容易，只要原数据有一点点改动，转换后结果就会有很大的差异。正是由于MD5算法的这些特性，它经常用于对于一段信息产生信息摘要，以防止其被篡改。其还广泛就于操作系统的登录过程中的安全验证，比如Unix操作系统的密码就是经过MD5加密后存储到文件系统中，当用户登录时输入密码后， 对用户输入的数据经过MD5加密后与原来存储的密文信息比对，如果相同说明密码正确，否则输入的密码就是错误的。
MD5以512位为一个计算单位对数据进行分组，每一分组又被划分为16个32位的小组，经过一系列处理后，输出4个32位的小组，最后组成一个128位的哈希值。对处理的数据进行512求余得到N和一个余数，如果余数不为448,填充1和若干个0直到448位为止，最后再加上一个64位用来保存数据的长度，这样经过预处理后，数据变成（N+1）x 512位。
加密。Encode 函数用来加密数据，Check函数传入一个未加密的字符串和与加密后的数据，进行对比，如果正确就返回true。

#### 5. Base64

Base64是一种任意二进制到文本字符串的编码方法，常用于在URL、Cookie、网页中传输少量二进制数据。
首先使用Base64编码需要一个含有64个字符的表，这个表由大小写字母、数字、+和/组成。采用Base64编码处理数据时，会把每三个字节共24位作为一个处理单元，再分为四组，每组6位，查表后获得相应的字符即编码后的字符串。编码后的字符串长32位，这样，经Base64编码后，原字符串增长1/3。如果要编码的数据不是3的倍数，最后会剩下一到两个字节，Base64编码中会采用\x00在处理单元后补全，编码后的字符串最后会加上一到两个 = 表示补了几个字节。

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
3. [密码学](https://blog.csdn.net/zhongliwen1981/article/details/104718931) 