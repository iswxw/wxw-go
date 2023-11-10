/*
@Time : 2022/4/18 16:02
@Author : weixiaowei
@File : demo_struct_copy
*/

package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
)

type AuthorInfo struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country *int   `json:"country"`
}

type Book struct {
	Title    string            `json:"title"`
	Author   AuthorInfo        `json:"author"`
	Year     int               `json:"year"`
	Category []string          `json:"category"`
	Price    map[string]string `json:"price"`
}

func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(&buffer).Decode(dst)
}

func DeepCopyByJson(src []Book) (*[]Book, error) {
	var dst = new([]Book)
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dst)
	return dst, err
}

func DeepCopyByCustom(src []Book) []Book {
	dst := make([]Book, len(src))
	for i, book := range src {
		tmpbook := Book{}
		tmpbook.Title = book.Title
		tmpbook.Year = book.Year
		tmpbook.Author = AuthorInfo{}
		tmpbook.Author.Name = book.Author.Name
		tmpbook.Author.Age = book.Author.Age
		tmpbook.Author.Country = new(int)
		*tmpbook.Author.Country = *book.Author.Country
		tmpbook.Category = make([]string, len(book.Category))
		for index, category := range book.Category {
			tmpbook.Category[index] = category
		}
		tmpbook.Price = make(map[string]string)
		for k, v := range book.Price {
			tmpbook.Price[k] = v
		}
		dst[i] = tmpbook
	}
	return dst
}

func check(err error) {
	if err != nil {
		panic(any(err))
	}
}

func printInfo(name string, books []Book) {
	for index, book := range books {
		fmt.Printf("%s[%d]=%v country=%d\n", name, index, book, *book.Author.Country)
	}
}

// 出自：https://blog.csdn.net/pengpengzhou/article/details/105839518
func main() {
	//初始化源Book切片
	books := make([]Book, 1)
	country := 1156
	author := AuthorInfo{"David", 38, &country}
	price := make(map[string]string)
	price["Europe"] = "$56"
	books[0] = Book{"Tutorial", author, 2020, []string{"math", "art"}, price}
	printInfo("books", books)

	var err error
	var start time.Time

	//Gob拷贝
	start = time.Now()
	booksCpy := make([]Book, 1)
	err = DeepCopyByGob(&booksCpy, books)
	fmt.Printf("\ngob time:%v\n", time.Now().Sub(start))
	check(err)
	*booksCpy[0].Author.Country = 1134
	booksCpy[0].Category[0] = "literature"
	booksCpy[0].Price["America"] = "$250"
	printInfo("booksCpy", booksCpy)
	printInfo("books", books)

	//JSON拷贝
	start = time.Now()
	booksCpy2, err_json := DeepCopyByJson(books)
	fmt.Printf("\njson time:%v\n", time.Now().Sub(start))
	check(err_json)
	*(*booksCpy2)[0].Author.Country = 1135
	(*booksCpy2)[0].Category[0] = "science"
	(*booksCpy2)[0].Price["Canada"] = "$150"
	printInfo("(*booksCpy2)", *booksCpy2)
	printInfo("books", books)

	//定制拷贝
	start = time.Now()
	booksCpy3 := DeepCopyByCustom(books)
	fmt.Printf("\ncustom time:%v\n", time.Now().Sub(start))

	*booksCpy3[0].Author.Country = 1136
	booksCpy3[0].Category[0] = "geometry"
	booksCpy3[0].Price["Africa"] = "$34"
	printInfo("booksCpy3", booksCpy3)
	printInfo("books", books)
}
