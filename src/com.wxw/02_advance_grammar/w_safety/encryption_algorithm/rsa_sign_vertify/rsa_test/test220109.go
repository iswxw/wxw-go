/*
@Time : 2022/1/9 22:13
@Author : wxw
@File : test
*/
package main

import (
	dd2201092 "com.wxw/AdvanceGrammar/w_safety/encryption_algorithm/rsa_sign_vertify/rsa_test/dd220109"
	"com.wxw/AdvanceGrammar/w_safety/encryption_algorithm/rsa_sign_vertify/utils"
	"fmt"
)

func main() {

	plains := "我们的人生不能靠心情活着,而要靠心态去生活."
	fmt.Println("加密前：" + plains)

	publicKey, _ := utils.ReadFileKey(utils.PemPath("public.pem"))
	cipherText, _ := dd2201092.EncryptBlock([]byte(plains), publicKey)
	fmt.Printf("加密后：%x\n", cipherText)

	privateKey, _ := utils.ReadFileKey(utils.PemPath("private.pem"))
	cipherPlains, _ := dd2201092.DecryptBlock(cipherText, privateKey)
	fmt.Println("解密后：" + string(cipherPlains))
}
