package tests

import (
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
