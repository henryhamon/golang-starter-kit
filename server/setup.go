package server

import (
	"golang-starter-kit/middleware"
	"golang-starter-kit/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.DBtoContext(db))
	router.InitializeRoutes(r)
	return r
}
