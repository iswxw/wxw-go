/*
@Time: 2021/9/21 11:24
@Author: wxw
@File: demo_01 数组的初始化
@Link: https://www.runoob.com/go/go-arrays.html
       https://www.liwenzhou.com/posts/Go/05_array/
*/
package main

import "fmt"

func main() {

	/*
	 * 【1】 每次都要确保提供的初始值和数组长度一致
	 */
	var nums1 [3]int                        // 数组会初始化为int类型的零值
	var nums2 = [3]int{1, 2, 3}             // 使用指定初始值完成初始化
	var nums3 = [3]string{"北京", "上海", "西安"} // 使用指定初始值完成初始化

	// 【2】 动态指定数组容量
	var nums4 = [...]string{"山东", "河南"}

	// 【3】使用指定索引值的方式初始化数组
	nums5 := [...]string{"泰山", "华山", "黄山"}

	fmt.Println(nums1) // [0 0 0]
	fmt.Println(nums2) // [1 2 3]
	fmt.Println(nums3) // [北京 上海 西安]

	fmt.Println(nums4)
	fmt.Printf("type of nums4:%T\n", nums4)

	fmt.Println(nums5)
	fmt.Printf("type of nums5:%T\n", nums5)

}
