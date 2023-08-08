// @Time : 2023/8/8 14:34
// @Author : xiaoweiwei
// @File : page_test

package demo_memory_page

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	// 假设这是我们的原始数据
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 设置分页参数
	pageSize := 3
	pageNumber := 2

	// 调用分页函数
	pagedData := paginate(data, pageNumber, pageSize)

	// 输出分页后的数据
	fmt.Println("Paged data:", pagedData)
}

// [出自]golang切片实现动态分页：https://www.jianshu.com/p/7898865e7c4c
// paginate 函数根据 pageNumber 和 pageSize 对原始数据进行分页
func paginate(data []int, pageNumber int, pageSize int) []int {
	// 获取原始数据长度
	dataLen := len(data)

	// 计算分页开始和结束的索引
	startIndex := (pageNumber - 1) * pageSize
	endIndex := startIndex + pageSize

	// 处理边界情况
	if startIndex > dataLen {
		return []int{}
	}

	if endIndex > dataLen {
		endIndex = dataLen
	}

	// 返回分页后的切片
	return data[startIndex:endIndex]
}
