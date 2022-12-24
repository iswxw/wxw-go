/*
@Time: 2022/12/24 23:13
@Author: wxw
@File: article
*/
package dao

// Tag 文章标签管理
type Tag struct {
	ID         uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `gorm:"column:name;type:varchar(100);comment:标签名称" json:"name"`
	CreatedOn  uint   `gorm:"column:created_on;type:int(11) unsigned;default:0;comment:创建时间" json:"created_on"`
	CreatedBy  string `gorm:"column:created_by;type:varchar(100);comment:创建人" json:"created_by"`
	ModifiedOn uint   `gorm:"column:modified_on;type:int(11) unsigned;default:0;comment:修改时间" json:"modified_on"`
	ModifiedBy string `gorm:"column:modified_by;type:varchar(100);comment:修改人" json:"modified_by"`
	DeletedOn  uint   `gorm:"column:deleted_on;type:int(11) unsigned;default:0;comment:删除时间" json:"deleted_on"`
	State      uint   `gorm:"column:state;type:tinyint(4) unsigned;default:1;comment:状态 0为禁用、1为启用" json:"state"`
}

// TableTag 表名
func (m *Tag) TableTag() string {
	return "blog_tag"
}
