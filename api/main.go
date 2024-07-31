package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	bootstrapping "github.com/milijan-mosic/NeoWarp/services/bootstrapping"
)

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
}

func main() {
	api := gin.Default()

	gin.SetMode(gin.ReleaseMode)
	api.ForwardedByClientIP = true
	api.SetTrustedProxies(nil)
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile)

	URL_PREFIX := "/api/1.0"
	bootstrapping.Bootstrap()

	// ROUTES

	api.GET(URL_PREFIX+"/", index)

	users := api.Group(URL_PREFIX + "/users")
	{
		users.GET("/all", bootstrapping.GetUsers)
		users.POST("/new", bootstrapping.CreateUser)
	}

	fmt.Println("Server listening on 4000...")
	api.Run(":4000")
}
