package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var err error

func main() {
	port, exist := os.LookupEnv("WEB_PORT")
	if !exist {
		port = "8080"
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Server run error:" + err.Error())
	}
}
