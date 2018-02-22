package main

import (
	"flag"
	db "golang-starter-kit/db"
	"golang-starter-kit/server"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	Port  string
	Debug string
)

func init() {
	Port = os.Getenv("PORT")
	Debug = os.Getenv("DEBUG")
	flag.StringVar(&Port, "p", Port, "Port")
	flag.StringVar(&Debug, "d", Debug, "DebugMode")
}

func main() {

	Db := db.Connect()
	defer Db.Close()

	if Debug == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := server.Setup(Db)

	port := "8080"

	if Port != "" {
		if _, err := strconv.Atoi(Port); err == nil {
			port = Port
		}
	}

	router.Run(":" + port)
}
