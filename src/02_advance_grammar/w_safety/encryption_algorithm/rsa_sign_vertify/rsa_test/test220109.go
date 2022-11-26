/*
@Time : 2022/1/9 22:13
@Author : wxw
@File : test
*/
package main

import (
	"fmt"
	dd2201092 "src/com.wxw/project_actual/src/02_advance_grammar/w_safety/encryption_algorithm/rsa_sign_vertify/rsa_test/dd220109"
	utils2 "src/com.wxw/project_actual/src/02_advance_grammar/w_safety/encryption_algorithm/rsa_sign_vertify/utils"
)

func main() {

	plains := "我们的人生不能靠心情活着,而要靠心态去生活."
	fmt.Println("加密前：" + plains)

	publicKey, _ := utils2.ReadFileKey(utils2.PemPath("public.pem"))
	cipherText, _ := dd2201092.EncryptBlock([]byte(plains), publicKey)
	fmt.Printf("加密后：%x\n", cipherText)

	privateKey, _ := utils2.ReadFileKey(utils2.PemPath("private.pem"))
	cipherPlains, _ := dd2201092.DecryptBlock(cipherText, privateKey)
	fmt.Println("解密后：" + string(cipherPlains))
}
