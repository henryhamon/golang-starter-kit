package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null;index:idx_user_nome" json:"name" form:"name" valid:"required"`
	Email    string `gorm:"not null;unique" json:"email" form:"email" valid:"email,required"`
	Username string `gorm:"not null;unique" json:"username" form:"username"`
	Pwd      string `json:"password" form:"pwd" valid:"length(6|20)"`
}
