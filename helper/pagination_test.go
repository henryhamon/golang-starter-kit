package helper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"

	"golang-starter-kit/models"

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
		db.DropTable(&models.User{})
		db.CreateTable(&models.User{})

		for i := 0; i < 5; i++ {
			user := models.User{}
			user.Name = fmt.Sprintf("%v", i)
			user.Email = fmt.Sprintf("%v@xyz.com", i)
			user.Username = fmt.Sprintf("%v", i)
			db.Create(&user)
		}

	}

	return db, err
}

func TestCountRecords(t *testing.T) {
	db, err := SetupDB()
	assert.NoError(t, err)

	users := []models.User{}
	done := make(chan bool, 1)
	var count int
	countRecords(db, &users, done, &count)
	assert.Equal(t, 5, count)
}

func TestTotalPages(t *testing.T) {
	assert.Equal(t, int64(6), getTotalPages(10, 60))
	assert.Equal(t, int64(6), getTotalPages(10, 59))
	assert.Equal(t, int64(1), getTotalPages(10, 3))
}

func TestPagination(t *testing.T) {
	db, err := SetupDB()
	assert.NoError(t, err)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/?limit=3", nil)

	params, err := NewParameter(c)
	assert.NoError(t, err)

	users := []models.User{}

	db, err = params.Paginate(db, &users)
	fmt.Println(users)
	assert.NoError(t, err)
	assert.Equal(t, 5, params.TotalRecords)
	assert.Equal(t, int64(2), params.TotalPages)

}
