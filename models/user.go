package models

import (
	"fmt"
	"golang-starter-kit/helper"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null;index:idx_user_nome" json:"name" form:"name" valid:"required"`
	Email    string `gorm:"not null;unique" json:"email" form:"email" valid:"email,required"`
	Username string `gorm:"not null;unique" json:"username" form:"username"`
	Pwd      string `json:"password" form:"pwd" valid:"length(6|20)"`
}

func ListUsers(db *gorm.DB, parameter *helper.Parameter) ([]User, error) {
	var err error
	users := []User{}

	for key, value := range parameter.Params {
		if value != "" {
			db = db.Where(fmt.Sprintf("%s LIKE ?", key), value+"%")
		}
	}

	rs, err := parameter.Paginate(db, &users)
	if err != nil {
		return []User{}, err
	}

	rs.Find(&users)

	return users, err
}
