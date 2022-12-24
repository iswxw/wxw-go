/*
@Time: 2022/12/24 23:17
@Author: wxw
@File: auth
*/
package dao

// Auth 授权信息
type Auth struct {
	Id       uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;type:varchar(50);comment:账号" json:"username"`
	Password string `gorm:"column:password;type:varchar(50);comment:密码" json:"password"`
}

func (m *Auth) TableAuth() string {
	return "blog_auth"
}
