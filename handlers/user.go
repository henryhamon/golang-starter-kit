package handlers

import (
	"net/http"

	dbpkg "golang-starter-kit/db"
	"golang-starter-kit/helper"
	"golang-starter-kit/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := dbpkg.DBInstance(c)
	parameter, err := helper.NewParameter(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	users, err := models.ListUsers(db, parameter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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
