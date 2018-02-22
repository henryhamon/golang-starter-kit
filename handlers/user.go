package handlers

import (
	"fmt"
	dbpkg "golang-starter-kit/db"
	"golang-starter-kit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := dbpkg.DBInstance(c)
	users := []models.User{}
	db.Find(&users)

	fmt.Println("aki")

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func NewUser(c *gin.Context) {

}

func GetUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
