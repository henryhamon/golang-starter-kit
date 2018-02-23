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

func GetUser(c *gin.Context) {
	db := dbpkg.DBInstance(c)
	user := models.User{}

	userId := c.Param("userId")
	if err := db.First(&user, userId).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	db := dbpkg.DBInstance(c)
	user := models.User{}
	userId := c.Param("userId")

	if db.First(&user, userId).RecordNotFound() {
		c.JSON(404, gin.H{"error": "user id " + userId + " not found!"})
		return
	}

	id := user.ID
	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.ID = id // TODO: Refactor
	if err := db.Save(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := dbpkg.DBInstance(c)
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, user)
}

func DeleteUser(c *gin.Context) {
	db := dbpkg.DBInstance(c)
	user := models.User{}
	userId := c.Param("userId")

	if db.First(&user, userId).RecordNotFound() {
		c.JSON(404, gin.H{"error": "user id " + userId + " not found!"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
