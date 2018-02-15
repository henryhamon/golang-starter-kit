package models

import (
	"testing"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	"github.com/stretchr/testify/assert"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupDB() (*gorm.DB, error) {

	var testconfig = struct {
		Test struct {
			Adapter  string `required:"true"`
			Database string `required:"true"`
		}
	}{}

	configor.Load(&testconfig, "../config.yml")

	db, err := gorm.Open(testconfig.Test.Adapter, testconfig.Test.Database)

	if err == nil {
		validations.RegisterCallbacks(db)
		db.DropTable(&User{})
		db.CreateTable(&User{})
	}

	return db, err
}

func TestValidations(t *testing.T) {
	db, err := SetupDB()
	if err != nil {
		t.Errorf("database connection error")
		return
	}
	defer db.Close()

	name := "Luke Skywalker"
	email := "luke@skywalker.com"
	pwd := "123456"

	user := User{}
	if err = db.Save(&user).Error; err == nil {
		t.Errorf("Must validate required fields")
	}

	assert.Contains(t, err.Error(), "blank")

	user.Name = name
	err = db.Save(&user).Error
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "email")
	}

	user.Email = "luke"
	err = db.Save(&user).Error
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "email address")
	}

	user.Email = email
	user.Pwd = "123"
	err = db.Save(&user).Error
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "length")
	}

	user.Pwd = pwd
	assert.NoError(t, db.Save(&user).Error)

	another := User{Name: name, Email: email, Pwd: pwd}
	err = db.Save(&another).Error
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Duplicate")
	}

}
