/*
@Time: 2022/1/9 14:29
@Author: wxw
@File: main
*/
package main

import (
	"fmt"
	utils "src/com.wxw/project_actual/src/com.wxw/project_actual/actual_safety/rsa_sign_vertify/utils"
)

func main() {

	// 生成密钥对，保存到文件路径： docs/tmp/
	utils.GeneratorRsaPairKeys(2048)

	// 测试加密和解密
	fmt.Println("============测试加密和解密==============")
	testDecryptAndEncrypt()

	// 签名和验签测试
	fmt.Println("============签名和验签测试==============")
	testSignAndVerify()

}

// 测试加密和解密
func testDecryptAndEncrypt() {
	// 待加密字符
	plains := "我们的人生不能靠心情活着,而要靠心态去生活,我们都不是完美的人,但要接受不完美的自己,学会独立,告别依赖,对软弱的自己说再见,永远不要停止相信自己!踏实一些,你想要的,岁月统统会还给你"
	// plains := "我们的人生不能靠心情活着,而要靠心态去生活."

	fmt.Println("加密前：" + plains)

	cipherText := utils.RSAEncrypt([]byte(plains), utils.PemPath("public.pem"))
	fmt.Printf("加密后：%x\n", cipherText)

	cipherPlains := utils.RSADecrypt(cipherText, utils.PemPath("private.pem"))
	fmt.Println("解密后：" + string(cipherPlains))
}

// 签名和验签测试
func testSignAndVerify() {
	plains := "我们的人生不能靠心情活着,而要靠心态去生活."

	fmt.Println("加密前：" + plains)

	cipherText := utils.RSAEncrypt([]byte(plains), utils.PemPath("public.pem"))
	fmt.Printf("加密后：%x\n", cipherText)

	cipherPlains := utils.RSADecrypt(cipherText, utils.PemPath("private.pem"))
	fmt.Println("解密后：" + string(cipherPlains))

	sign := utils.Sign([]byte(plains), utils.PemPath("private.pem"))
	fmt.Printf("签名时：%x\n", sign)

	// plains = "我们的人生不能靠心情活着,而要靠心态去生活"
	verify := utils.Verify([]byte(plains), utils.PemPath("public.pem"), sign)
	fmt.Printf("验签时: %t\n", verify)

}
