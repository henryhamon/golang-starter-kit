package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"

	"golang-starter-kit/models"
	"golang-starter-kit/tester"
)

func beforeStart(db *gorm.DB) []models.User {
	db.DropTable(&models.User{})
	db.CreateTable(&models.User{})
	users := []models.User{}

	for i := 0; i < 2; i++ {
		user := models.User{}
		user.Name = fmt.Sprintf("%v", i)
		user.Email = fmt.Sprintf("%v@xyz.com", i)
		user.Username = fmt.Sprintf("%v", i)
		db.Create(&user)
		users = append(users, user)
	}

	return users
}

func TestGetUserList(t *testing.T) {
	db, err := tester.SetupDB()
	defer db.Close()
	assert.NoError(t, err)
	users := beforeStart(db)
	userResp := []models.User{}

	w := tester.PerformRequest(tester.GinEngine(db), "GET", "/v1/users")

	var raw map[string][]models.User
	err = json.Unmarshal([]byte(w.Body.String()), &raw)
	userResp = raw["users"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(users), len(userResp))
	assert.Equal(t, users[0].Name, userResp[0].Name)
}

func TestGetUser(t *testing.T) {
	db, err := tester.SetupDB()
	defer db.Close()
	assert.NoError(t, err)
	users := beforeStart(db)
	userResp := models.User{}
	var raw map[string]models.User

	w := tester.PerformRequest(tester.GinEngine(db), "GET", "/v1/user/show/1")
	err = json.Unmarshal([]byte(w.Body.String()), &raw)
	assert.NoError(t, err)

	userResp = raw["user"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, users[0].Name, userResp.Name)
}

func TestDeleteUser(t *testing.T) {
	var count int64
	db, err := tester.SetupDB()
	defer db.Close()
	assert.NoError(t, err)
	beforeStart(db)

	db.Where("deleted_at is null").Table("users").Count(&count)
	assert.Equal(t, int64(2), count)
	tester.PerformRequest(tester.GinEngine(db), "DELETE", "/v1/user/1")

	db.Where("deleted_at is null").Table("users").Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestInsertUser(t *testing.T) {
	var count int64
	db, err := tester.SetupDB()
	defer db.Close()
	assert.NoError(t, err)
	beforeStart(db)

	db.Table("users").Count(&count)
	assert.Equal(t, int64(2), count)

	req, err := json.Marshal(models.User{Name: "test", Email: "test@test.com", Username: "tester"})
	assert.NoError(t, err)
	w := tester.PerformRequest(tester.GinEngine(db), "POST", "/v1/user/new", bytes.NewReader(req))

	assert.Equal(t, 201, w.Code)

	db.Table("users").Count(&count)
	assert.Equal(t, int64(3), count)
}

func TestUpdateUser(t *testing.T) {
	db, err := tester.SetupDB()
	defer db.Close()
	assert.NoError(t, err)
	beforeStart(db)

	req, err := json.Marshal(models.User{Name: "test", Email: "test@test.com", Username: "tester"})
	assert.NoError(t, err)
	w := tester.PerformRequest(tester.GinEngine(db), "PUT", "/v1/user/1", bytes.NewReader(req))

	assert.Equal(t, 200, w.Code)

	user := models.User{}
	db.First(&user, 1)
	assert.Equal(t, "test", user.Name)
}
