/*
@Time: 2022/9/5 22:31
@Author: wxw  https://zhuanlan.zhihu.com/p/482547957
@File: demo01_advance
*/
package demo01_reflect

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

// DeleteSliceElms 从切片中过滤指定元素。注意：不修改原切片。
func DeleteSliceElms(i interface{}, elms ...interface{}) interface{} {
	// 构建 map set。
	m := make(map[interface{}]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	v := reflect.ValueOf(i)                               // 取值
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len()) // 构造切片
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface()
}

// DeleteU64liceElms 从 []uint64 过滤指定元素。注意：不修改原切片。
func DeleteU64liceElms(i []uint64, elms ...uint64) []uint64 {
	// 构建 map set。
	m := make(map[uint64]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	t := make([]uint64, 0, len(i))
	for _, v := range i {
		if _, ok := m[v]; !ok {
			t = append(t, v)
		}
	}
	return t
}

// Ntohl 将网络字节序的 uint32 转为主机字节序。
func Ntohl(bys []byte) uint32 {
	buf := bytes.NewReader(bys)
	var num uint32
	binary.Read(buf, binary.BigEndian, &num)
	return num
}

func NtohlNotUseBinary(bys []byte) uint32 {
	return uint32(bys[3]) | uint32(bys[2])<<8 | uint32(bys[1])<<16 | uint32(bys[0])<<24
}
