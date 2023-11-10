/*
@Time: 2022/12/24 23:19
@Author: wxw
@File: article
*/
package dao

// Article 文章管理
type Article struct {
	Id            uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	TagId         uint   `gorm:"column:tag_id;type:int(11) unsigned;default:0;comment:标签ID" json:"tag_id"`
	Title         string `gorm:"column:title;type:varchar(100);comment:文章标题" json:"title"`
	Desc          string `gorm:"column:desc;type:varchar(255);comment:简述" json:"desc"`
	Content       string `gorm:"column:content;type:text;comment:内容" json:"content"`
	CoverImageUrl string `gorm:"column:cover_image_url;type:varchar(255);comment:封面图片地址" json:"cover_image_url"`
	CreatedOn     uint   `gorm:"column:created_on;type:int(11) unsigned;default:0;comment:新建时间" json:"created_on"`
	CreatedBy     string `gorm:"column:created_by;type:varchar(100);comment:创建人" json:"created_by"`
	ModifiedOn    uint   `gorm:"column:modified_on;type:int(11) unsigned;default:0;comment:修改时间" json:"modified_on"`
	ModifiedBy    string `gorm:"column:modified_by;type:varchar(255);comment:修改人" json:"modified_by"`
	DeletedOn     uint   `gorm:"column:deleted_on;type:int(11) unsigned;default:0" json:"deleted_on"`
	State         uint   `gorm:"column:state;type:tinyint(4) unsigned;default:1;comment:删除时间" json:"state"`
}

func (m *Article) TableArticle() string {
	return "blog_article"
}
