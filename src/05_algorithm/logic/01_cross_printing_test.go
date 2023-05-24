// @Time : 2023/5/24 16:11
// @Author : xiaoweiwei 数据和字母交叉打印
// @File : 01_cross_printing_test

package logic

import (
	"fmt"
	"sync"
	"testing"
)

func TestCrossPrinting(t *testing.T) {

	var wg sync.WaitGroup
	letter, number := make(chan string, 26), make(chan int, 27)
	wg.Add(3)

	// 记录字母
	go func() {
		defer wg.Done()
		for i := 'A'; i <= 'Z'; i++ {
			letter <- string(i)
		}
		close(letter)
	}()

	// 记录数字
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			number <- i
		}
		close(number)
	}()

	// 打印数字和字母
	go func() {
		defer wg.Done()
		for i := range number {
			fmt.Printf("%d%s", i, <-letter)
		}
	}()
	wg.Wait()

	fmt.Println("\n", "")

}

// 关于字符和整型的说明：https://blog.csdn.net/wucz122140729/article/details/105724633
// ASCII码 对照表：http://ascii.wjccx.com/
// TestCharAndIntContact 字符与整型之间的关系
func TestCharAndIntContact(t *testing.T) {
	fmt.Println(fmt.Sprintf("%c", 'A'+0))
	fmt.Println(fmt.Sprintf("%c", 'A'+1))
	fmt.Println("—————————————————————————————")
	fmt.Println('A', string('A'))
	fmt.Println(rune('A'), int32('a'))
	fmt.Println(rune('A' + 1))
	fmt.Println(string(rune('A' + 1)))
}
