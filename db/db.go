package db

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

	"golang-starter-kit/models"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Config is loaded from configuration file
var Config = struct {
	Test struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}

	Production struct {
		Adapter  string `required:"true"`
		Database string `required:"true"`
	}
}{}

var (
	Migrate string
	ShowSql string
)

func init() {
	Migrate = os.Getenv("MIGRATE")
	flag.StringVar(&Migrate, "m", Migrate, "AutoMigrate")
	flag.StringVar(&ShowSql, "s", ShowSql, "ShowSql")
}

func Connect() *gorm.DB {
	flag.Parse()
	configor.Load(&Config, "config.yml")

	db, err := gorm.Open(Config.Production.Adapter, Config.Production.Database)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	validations.RegisterCallbacks(db)

	db.LogMode(false)

	if ShowSql != "" {
		db.LogMode(true)
	}

	if Migrate == "1" {
		db.AutoMigrate(
			&models.User{},
		)

	}
	return db
}

func DBInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}
