package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"golang-starter-kit/helper"
	"golang-starter-kit/models"
	"golang-starter-kit/tester"
)

func createUsersBeforeStart(db *gorm.DB) {
	db.DropTable(&models.User{})
	db.CreateTable(&models.User{})
	users := []models.User{}

	for i := 0; i < 5; i++ {
		user := models.User{}
		user.Name = fmt.Sprintf("%v", i)
		user.Email = fmt.Sprintf("%v@xyz.com", i)
		user.Username = fmt.Sprintf("%v", i)
		db.Create(&user)
		users = append(users, user)
	}

}

func TestCountRecords(t *testing.T) {
	db, err := tester.SetupDB()
	assert.NoError(t, err)
	createUsersBeforeStart(db)

	users := []models.User{}
	done := make(chan bool, 1)
	var count int
	helper.CountRecords(db, &users, done, &count)
	assert.Equal(t, 5, count)
}

func TestTotalPages(t *testing.T) {
	assert.Equal(t, int64(6), helper.GetTotalPages(10, 60))
	assert.Equal(t, int64(6), helper.GetTotalPages(10, 59))
	assert.Equal(t, int64(1), helper.GetTotalPages(10, 3))
}

func TestPagination(t *testing.T) {
	db, err := tester.SetupDB()
	assert.NoError(t, err)
	createUsersBeforeStart(db)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "http://test.com/?limit=3", nil)

	params, err := helper.NewParameter(c)
	assert.NoError(t, err)

	users := []models.User{}

	db, err = params.Paginate(db, &users)
	fmt.Println(users)
	assert.NoError(t, err)
	assert.Equal(t, 5, params.TotalRecords)
	assert.Equal(t, int64(2), params.TotalPages)

}
