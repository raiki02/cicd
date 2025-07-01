package main

import (
	"github.com/gin-gonic/gin"
)

func newApp() *gin.Engine {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.Logger())
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return app
}
