package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/", func(c *gin.Context) {
		dateTime := time.Now()
		c.String(http.StatusOK, dateTime.String())
	})

	return router
}

func main() {
	router := setUpRouter()
	port := os.Getenv("PORT")

	if port != "" {
		router.Run("0.0.0.0:" + port)
	} else {
		router.Run("0.0.0.0:5001")
	}

}
