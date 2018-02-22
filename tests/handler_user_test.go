package tests

import (
	"testing"
)

func TestGetUserList(t *testing.T) {
	// r := gofight.New()
	/*
		db, err := tester.SetupDB()
		defer db.Close()
		assert.NoError(t, err)
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

		req, err := http.NewRequest("GET", "/v1/users/", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		tester.GinEngine(db).ServeHTTP(resp, req)
		assert.Equal(t, 201, resp.Code)
		// */

	// r.GET("/v1/users/").
	// 	SetDebug(true).
	// 	Run(tester.GinEngine(db), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
	// 		fmt.Println(r.Body.String())

	// 		assert.Equal(t, http.StatusOK, r.Code)
	// 		assert.Equal(t, "application/json; charset=utf-8", r.HeaderMap.Get("Content-Type"))
	// 	})
}
