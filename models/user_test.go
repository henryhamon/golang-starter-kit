package models

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-starter-kit/helper"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

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

func TestListUsers(t *testing.T) {
	db, err := SetupDB()
	assert.NoError(t, err)

	for i := 0; i < 5; i++ {
		user := User{}
		user.Name = fmt.Sprintf("%v", i)
		user.Email = fmt.Sprintf("%v@xyz.com", i)
		user.Username = fmt.Sprintf("%v", i)
		db.Create(&user)
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/?limit=3", nil)

	params, err := helper.NewParameter(c)
	assert.NoError(t, err)

	users, err := ListUsers(db, params)
	assert.Equal(t, 3, len(users))
}

func TestListOnlyOneUser(t *testing.T) {
	db, err := SetupDB()
	assert.NoError(t, err)

	for i := 0; i < 5; i++ {
		user := User{}
		user.Name = fmt.Sprintf("%v", i)
		user.Email = fmt.Sprintf("%v@xyz.com", i)
		user.Username = fmt.Sprintf("%v", i)
		db.Create(&user)
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/?limit=3&name=1", nil)

	params, err := helper.NewParameter(c)
	assert.NoError(t, err)

	users, err := ListUsers(db, params)
	assert.Equal(t, 1, len(users))
}
