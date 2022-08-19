package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	timer "github.com/HendricksK/timer-service/timer"
	"github.com/gin-gonic/gin"
)

var port string
var env string

func setUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/", func(c *gin.Context) {
		dateTime := time.Now()
		c.String(http.StatusOK, fmt.Sprintf("%v\n%v\n", dateTime.String(), "https://www.youtube.com/watch?v=HTFmOOwhdd4"))
	})

	// Timer CRUD
	// This is a baseline test URI
	router.GET("/timers", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, timer.Read())
	})

	router.GET("/timer/:ref", func(c *gin.Context) {
		ref := c.Param("ref")
		fmt.Println(ref)
		c.IndentedJSON(http.StatusOK, timer.ReadById(ref))
	})

	router.POST("/timer", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, timer.Create(c))
	})

	router.PATCH("/timer/:ref", func(c *gin.Context) {
		ref := c.Param("ref")
		fmt.Println(ref)
		c.IndentedJSON(http.StatusOK, timer.Update(ref, c))
	})

	router.DELETE("/timer/:ref", func(c *gin.Context) {
		ref := c.Param("ref")
		fmt.Println(ref)
		c.IndentedJSON(http.StatusOK, timer.Delete(ref))
	})

	return router
}

func main() {
	router := setUpRouter()
	// port := os.Getenv("PORT")
	// env := os.Getenv("ENV")

	fmt.Println(env)
	fmt.Println(port)

	if port != "" {
		router.Run("0.0.0.0:" + port)
	} else {
		router.Run("0.0.0.0:5001")
	}
}

// https://tutorialedge.net/golang/the-go-init-function/
func init() {
	port = os.Getenv("PORT")
	env = os.Getenv("ENV")
	fmt.Println(timer.Init())
}
