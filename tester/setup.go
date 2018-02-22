package tester

import (
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
	}

	return db, err
}

func GinEngine(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	return r
}
