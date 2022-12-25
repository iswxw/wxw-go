/*
@Time: 2022/12/24 23:22
@Author: wxw
@File: article_model
*/
package model

import (
	"gorm.io/gorm"
	"src/com.wxw/project_actual/module/gin-example/app/blog/dao"
	"src/com.wxw/project_actual/module/gin-example/common/infra/mysql"
)

type IArticle interface {
	Query(tagID uint) (*dao.Article, error)
	Insert(article *dao.Article) (uint, error)
}

type articleModel struct {
	db *gorm.DB
}

func (a articleModel) Insert(article *dao.Article) (uint, error) {
	db := a.db.Table(article.TableArticle())
	err := db.Create(&article).Error
	if err != nil {
		return 0, err
	}
	return article.TagId, nil
}

func (a articleModel) Query(tagID uint) (*dao.Article, error) {
	article := &dao.Article{}
	db := a.db.Table(article.TableArticle())

	if tagID != 0 {
		db = db.Where("tag_id = ?", tagID)
	}

	err := db.Order("id desc").First(article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

// NewArticleModel 初始化一个实例
func NewArticleModel() IArticle {
	return &articleModel{
		db: mysql.Client,
	}
}
