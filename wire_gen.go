package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func initDB() {
	dsn := "root:root@tcp(192.168.43.194:3306)/cicd?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
}

func newApp() *gin.Engine {
	// initDB()
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(gin.Logger())
	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":     "pong",
			"DBconnected": db != nil,
		})
	})
	return app
}
