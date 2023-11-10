/*
@Time: 2022/12/25 11:39
@Author: wxw
@File: article_model_test
*/
package model

import (
	"github.com/spf13/cast"
	"log"
	"src/com.wxw/project_actual/module/gin-example/app/blog/dao"
	"src/com.wxw/project_actual/module/gin-example/test"
	"testing"
	"time"
)

func TestArticleModel_Insert(t *testing.T) {
	test.SetupTest()
	insert, err := NewArticleModel().Insert(&dao.Article{
		TagId:         10000,
		Title:         "技术能量站",
		Desc:          "博客主页",
		Content:       "博客内容",
		State:         1,
		CoverImageUrl: "封面图片",
		CreatedOn:     cast.ToUint(time.Now().Unix()),
		CreatedBy:     "wxw",
	})
	log.Println("err = ", err)
	log.Println("insert = ", insert)
}

func TestArticleModel_Query(t *testing.T) {
	test.SetupTest()
	article, err := NewArticleModel().Query(10000)
	log.Println("err = ", err)
	log.Println("article = ", article)
}
