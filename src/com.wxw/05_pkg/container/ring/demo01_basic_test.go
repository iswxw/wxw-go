package main

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {

	fmt.Println("测试空链表")
	var rings ring.Ring
	fmt.Println(rings.Next().Value)

	fmt.Println("")
	fmt.Println("")

	fmt.Println("测试环")
	newRing := makeN(10)
	Print(newRing)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("测试环Link自己")
	newRing = makeN(10)
	newRing = newRing.Link(newRing)
	Print(newRing)

	fmt.Println()
	fmt.Println()
	fmt.Println("测试环Link其他环")
	newRing = makeN(10)
	newRing1 := makeN(5)
	newRing = newRing.Link(newRing1)
	Print(newRing)

	fmt.Println()
	fmt.Println()
	fmt.Println("测试Move")
	newRing1 = makeN(5)
	newRing1 = newRing1.Move(3)
	Print(newRing1)

	fmt.Println()
	fmt.Println()
	fmt.Println("测试Do，没有办法改变环中的元素.")
	newRing1 = makeN(5)
	newRing1.Do(func(i interface{}) {
		i = i.(int) + 1000
	})
	Print(newRing1)

	newRingx := ring.New(3)

	for i := 1; i <= newRingx.Len(); i++ {
		newRingx.Value = Person{
			Name:  "小明",
			Age:   100,
			Money: 19999,
		}
		newRingx = newRingx.Next()
		newRingx.Value = Person{
			Name:  "小康",
			Age:   50,
			Money: 19921999,
		}
		newRingx = newRingx.Next()
		newRingx.Value = Person{
			Name:  "小施",
			Age:   20,
			Money: 1111999,
		}
	}
	fmt.Println("测试Do，只能对元素进行提取或者输出.")
	s := make([]int, 0)
	newRingx.Do(func(i interface{}) {
		s = append(s, i.(Person).Age)
	})
	for _, item := range s {
		fmt.Println(item)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("测试Unlink")
	newRing1 = makeN(5)
	newRing2 := newRing1.Unlink(2)
	Print(newRing1)
	println()
	fmt.Println("如果用赋值语句，可以返回被移除的元素")
	Print(newRing2)
}

func Print(r *ring.Ring) {
	i, n := 0, r.Len()
	for p := r; i < n; p = p.Next() {
		fmt.Println(fmt.Sprintf("当前元素是%v", p.Value))
		i++
	}
}

//创造对应的链表
func makeN(n int) *ring.Ring {
	r := ring.New(n)
	for i := 1; i <= n; i++ {
		r.Value = i
		r = r.Next()
	}
	return r
}

type Person struct {
	Name  string
	Age   int
	Money int
}
