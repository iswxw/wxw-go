package main

import (
	"container/list"
	"fmt"
	"testing"
)

// list 结构基础测试
func TestList(t *testing.T) {
	list2 := list.New()

	list2.PushFront(Person{
		Name:  "康康",
		Age:   10,
		Money: 20,
	})

	list2.PushBack(Person{
		Name:  "老施",
		Age:   40,
		Money: 1000000,
	})

	for e := list2.Front(); e != nil; e = e.Next() {
		value := e.Value.(Person)
		fmt.Println(fmt.Sprintf("名字：%v,年龄：%v，身家:%v", value.Name, value.Age, value.Money))
	}
	fmt.Println()
	fmt.Println()

	frontE := list2.Front()
	list2.MoveToBack(frontE)

	for e := list2.Front(); e != nil; e = e.Next() {
		value := e.Value.(Person)
		fmt.Println(fmt.Sprintf("名字：%v,年龄：%v，身家:%v", value.Name, value.Age, value.Money))
	}
	fmt.Println()
	fmt.Println()

	xiaozhang := list2.InsertBefore(Person{
		Name:  "小张",
		Age:   10,
		Money: 50,
	}, list2.Front())
	for e := list2.Front(); e != nil; e = e.Next() {
		value := e.Value.(Person)
		fmt.Println(fmt.Sprintf("名字：%v,年龄：%v，身家:%v", value.Name, value.Age, value.Money))
	}
	fmt.Println()
	fmt.Println()

	list2.Remove(xiaozhang)
	for e := list2.Front(); e != nil; e = e.Next() {
		value := e.Value.(Person)
		fmt.Println(fmt.Sprintf("名字：%v,年龄：%v，身家:%v", value.Name, value.Age, value.Money))
	}
	fmt.Println()
	fmt.Println()

	front := list2.Front().Value.(Person)
	fmt.Println(fmt.Sprintf("名字：%v,年龄：%v，身家:%v", front.Name, front.Age, front.Money))
}

type Person struct {
	Name  string
	Age   int
	Money int
}
