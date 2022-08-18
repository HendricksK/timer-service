package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	databaseconnector "github.com/HendricksK/timer-service/database-connector"
	timer "github.com/HendricksK/timer-service/timer"
	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/", func(c *gin.Context) {
		dateTime := time.Now()
		c.String(http.StatusOK, fmt.Sprintf("%v\n%v\n", dateTime.String(), "https://www.youtube.com/watch?v=HTFmOOwhdd4"))
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

// https://tutorialedge.net/golang/the-go-init-function/
func init() {
	fmt.Println(databaseconnector.Init())
	fmt.Println(timer.Init())
}
